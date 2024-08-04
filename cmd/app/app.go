package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/pathfinder"
)

type App struct {
	config *config.Config
	goPath string
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
		goPath: pathfinder.GetGoPath(),
	}
}

func (a *App) Run(_ context.Context) int {
	args := os.Args[1:]
	linterOutput, _ := exec.Command(a.goPath+a.config.InstallDir+"/golangci-lint", args...).Output()
	fmt.Println(string(linterOutput))

	return exitcode.Success
}
