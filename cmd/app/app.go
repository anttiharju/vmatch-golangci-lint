package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/exitcode"
)

type app struct{}

func NewApp() *app {
	return &app{}
}

func (a *app) Run(_ context.Context) int {
	args := os.Args[1:]

	goBinPath, _ := exec.LookPath("go")
	// fmt.Println(goBinPath)

	goPathBytes, _ := exec.Command(goBinPath, "env", "GOPATH").Output()
	goPath := string(goPathBytes)
	goPath = strings.TrimSpace(goPath)
	// fmt.Println(goPath)

	linterBinPath := goPath + "/bin/golangci-lint"
	// fmt.Println(linterBinPath)

	out, _ := exec.Command(linterBinPath, args...).Output()
	fmt.Println(string(out))

	return exitcode.Success
}
