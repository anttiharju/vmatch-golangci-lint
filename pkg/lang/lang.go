package lang

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch/pkg/exitcode"
	"github.com/anttiharju/vmatch/pkg/finder"
	"github.com/anttiharju/vmatch/pkg/wrapper"
)

type WrappedLang struct {
	wrapper.BaseWrapper
	desiredVersion string
}

func NewWrapper(name string) *WrappedLang {
	baseWrapper := wrapper.BaseWrapper{Name: name}

	desiredVersion, err := finder.GetLangVersion()
	if err != nil {
		baseWrapper.ExitWithPrintln(exitcode.VersionReadFileIssue, err.Error())
	}

	err = baseWrapper.GenerateInstallPath(desiredVersion)
	if err != nil {
		baseWrapper.ExitWithPrintln(exitcode.InstallPathIssue, err.Error())
	}

	return &WrappedLang{
		BaseWrapper:    baseWrapper,
		desiredVersion: desiredVersion,
	}
}

func (w *WrappedLang) Run(ctx context.Context, args []string) int {
	if w.noBinary() {
		w.install(ctx)
	}

	//nolint:gosec // I don't think a wrapper can avoid G204.
	lang := exec.Command(w.getGolangCILintPath(), args...)
	langOutput, _ := lang.Output()

	fmt.Print(string(langOutput))

	return lang.ProcessState.ExitCode()
}

func (w *WrappedLang) noBinary() bool {
	_, err := os.Stat(w.getGolangCILintPath())

	return os.IsNotExist(err)
}

func (w *WrappedLang) install(_ context.Context) {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	// todo: pin to a sha instead of master, but automate updates
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	command := curl + pipe + sh + w.InstallPath + " v" + w.desiredVersion
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

func (w *WrappedLang) getGolangCILintPath() string {
	return w.InstallPath + string(os.PathSeparator) + "golangci-lint"
}

var _ wrapper.Interface = (*WrappedLang)(nil)
