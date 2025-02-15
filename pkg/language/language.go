package language

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anttiharju/vmatch/pkg/exitcode"
	"github.com/anttiharju/vmatch/pkg/finder"
	"github.com/anttiharju/vmatch/pkg/wrapper"
)

type WrappedLanguage struct {
	wrapper.BaseWrapper
}

func langParser(content []byte) (string, error) {
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			trimmed := strings.TrimPrefix(line, "go ")

			return trimmed, nil
		}
	}

	return "", errors.New("cannot find go version")
}

func NewWrapper(name string) *WrappedLanguage {
	baseWrapper := wrapper.BaseWrapper{Name: name}

	desiredVersion, err := finder.GetVersion("go.mod", langParser)
	if err != nil {
		baseWrapper.ExitWithPrintln(exitcode.VersionReadFileIssue, err.Error())
	}

	err = baseWrapper.GenerateInstallPath(desiredVersion)
	if err != nil {
		baseWrapper.ExitWithPrintln(exitcode.InstallPathIssue, err.Error())
	}

	return &WrappedLanguage{
		BaseWrapper: baseWrapper,
	}
}

func (w *WrappedLanguage) Run(args []string) int {
	if w.noBinary() {
		w.install()
	}

	//nolint:gosec // I don't think a wrapper can avoid G204.
	lang := exec.Command(w.getGolangCILintPath(), args...)
	langOutput, _ := lang.Output()

	fmt.Print(string(langOutput))

	return lang.ProcessState.ExitCode()
}

func (w *WrappedLanguage) noBinary() bool {
	_, err := os.Stat(w.getGolangCILintPath())

	return os.IsNotExist(err)
}

func (w *WrappedLanguage) install() {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	// todo: pin to a sha instead of master, but automate updates
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	command := curl + pipe + sh + w.InstallPath + " v" + w.DesiredVersion
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		w.ExitWithPrint(exitcode.CmdStartIssue, "failed to start command: "+err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		w.ExitWithPrint(exitcode.CmdStartIssue, "failed to wait for command: "+err.Error())
	}
}

func (w *WrappedLanguage) getGolangCILintPath() string {
	return w.InstallPath + string(os.PathSeparator) + "golangci-lint"
}

var _ wrapper.Interface = (*WrappedLanguage)(nil)
