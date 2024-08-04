package app

import (
	"context"
	"fmt"
	"os"

	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/exitcode"
)

type app struct{}

func NewApp() *app {
	return &app{}
}

func (a *app) Run(_ context.Context) int {
	args := os.Args[1:]

	fmt.Println("Received args:")

	for i, arg := range args {
		fmt.Printf("%d %s\n", i+1, arg)
	}

	return exitcode.Success
}
