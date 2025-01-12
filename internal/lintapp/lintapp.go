package lintapp

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch/pkg/app"
	"github.com/anttiharju/vmatch/pkg/exit"
	"github.com/anttiharju/vmatch/pkg/exit/exitcode"
	"github.com/anttiharju/vmatch/pkg/pathfinder"
	"github.com/anttiharju/vmatch/pkg/versionfinder"
)

type LintApp struct {
	desiredVersion string
	installPath    string
}

var _ app.Interface = (*LintApp)(nil)

func NewApp(versionFileName string) *LintApp {
	desiredVersion := versionfinder.GetVersion(pathfinder.GetWorkDir(), versionFileName)
	installPath := pathfinder.GetInstallPath(desiredVersion)

	return &LintApp{
		desiredVersion: desiredVersion,
		installPath:    installPath,
	}
}

func (l *LintApp) Run(ctx context.Context) int {
	if l.noBinary() {
		l.install(ctx)
	}

	args := os.Args[1:]
	//nolint:gosec // I don't think a wrapper can avoid G204.
	linter := exec.Command(l.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (l *LintApp) noBinary() bool {
	_, err := os.Stat(l.getGolangCILintPath())

	return os.IsNotExist(err)
}

func (l *LintApp) install(_ context.Context) {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	command := curl + pipe + sh + l.installPath + " v" + l.desiredVersion
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		exit.WithMessage(exitcode.CmdStartIssue, "failed to start command: "+err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		exit.WithMessage(exitcode.CmdStartIssue, "failed to wait for command: "+err.Error())
	}
}

func (l *LintApp) getGolangCILintPath() string {
	return l.installPath + string(os.PathSeparator) + "golangci-lint"
}
