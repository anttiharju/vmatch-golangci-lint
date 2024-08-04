package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/exitcode"
	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/pathfinder"
)

type app struct {
	goPath string
}

func NewApp() *app {
	return &app{goPath: pathfinder.GetGoPath()}
}

func (a *app) Run(_ context.Context) int {
	args := os.Args[1:]

	out, _ := exec.Command(a.goPath+"/bin/golangci-lint", args...).Output()
	fmt.Println(string(out))

	return exitcode.Success
}
