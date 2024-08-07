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
	goPath         string
	desiredVersion string
}

func NewApp(config *config.Config) *App {
	return &App{
		config:         config,
		goPath:         pathfinder.GetGoPath(),
		desiredVersion: versionfinder.GetVersion(config.VersionFileName),
	}
}

func (a *App) Run(_ context.Context) int {
	args := os.Args[1:]
	linter := exec.Command(a.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (a *App) getGolangCILintPath() string {
	return a.goPath + a.config.InstallDir + string(os.PathSeparator) + "golangci-lint"
}
