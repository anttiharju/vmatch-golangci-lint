package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/homebrew-golangci-lint-updater/src/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/pathfinder"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/versionfinder"
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

	fmt.Println(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (a *App) getGolangCILintPath() string {
	return a.goPath + a.config.InstallDir + string(os.PathSeparator) + "golangci-lint"
}
