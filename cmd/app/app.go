package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/exitcode"
	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/pathfinder"
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

	out, _ := exec.Command(a.goPath+a.config.InstallDir+"/golangci-lint", args...).Output()
	fmt.Println(string(out))

	return exitcode.Success
}
