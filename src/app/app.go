package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch-golangci-lint/src/finder"
)

type App struct {
	desiredVersion string
	installPath    string
}

func NewApp(versionFileName string) *App {
	desiredVersion := finder.GetVersion(versionFileName)
	installPath := finder.GetBinDir() + string(os.PathSeparator) + desiredVersion

	return &App{
		desiredVersion: desiredVersion,
		installPath:    installPath,
	}
}

func (a *App) Run(ctx context.Context) int {
	if a.needToDownload() {
		a.install(ctx)
	}

	args := os.Args[1:]
	linter := exec.Command(a.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (a *App) needToDownload() bool {
	_, err := os.Stat(a.getGolangCILintPath())
	return os.IsNotExist(err)
}

func (a *App) install(_ context.Context) {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	command := curl + pipe + sh + a.installPath + " " + a.desiredVersion
	cmd := exec.Command("sh", "-c", command)
	cmd.Start()
	cmd.Wait()
}

func (a *App) getGolangCILintPath() string {
	return a.installPath + string(os.PathSeparator) + "golangci-lint"
}
