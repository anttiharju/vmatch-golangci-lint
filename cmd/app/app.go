package app

import (
	"context"

	"github.com/anttiharju/homebrew-golangci-lint-updater/pkg/exitcode"
)

type app struct{}

func NewApp() *app {
	return &app{}
}

func (a *app) Run(_ context.Context) int {
	return exitcode.Success
}
