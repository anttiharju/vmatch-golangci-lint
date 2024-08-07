package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch-golangci-lint/src/config"
	"github.com/anttiharju/vmatch-golangci-lint/src/pathfinder"
	"github.com/anttiharju/vmatch-golangci-lint/src/versionfinder"
)

type App struct {
	config         *config.Config
	desiredVersion string
	installPath    string
}

func NewApp(config *config.Config) *App {
	desiredVersion := versionfinder.GetVersion(config.VersionFileName)
	ps := string(os.PathSeparator)
	v := string(desiredVersion[0])
	numbers := desiredVersion[1:]
	installPath := pathfinder.GetBinDir() + ps + v + ps + numbers

	return &App{
		config:         config,
		desiredVersion: desiredVersion,
		installPath:    installPath,
	}
}

func (a *App) Run(ctx context.Context) int {
	a.install(ctx)

	args := os.Args[1:]
	linter := exec.Command(a.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (a *App) install(_ context.Context) {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	cmd := curl + pipe + sh + a.installPath + " " + a.desiredVersion
	execCmd := exec.Command("sh", "-c", cmd)
	execCmd.Start()
	execCmd.Wait()
}

func (a *App) getGolangCILintPath() string {
	return a.installPath + string(os.PathSeparator) + "golangci-lint"
}
