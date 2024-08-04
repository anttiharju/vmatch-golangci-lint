package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/app"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exit"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
)

func main() {
	ctx := context.Background()

	go listenInterrupts()

	config := config.NewConfig()
	app := app.NewApp(config)
	exitCode := app.Run(ctx)
	exit.Now(exitCode)
}

func listenInterrupts() {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	<-interruptCh
	exit.WithMessage(exitcode.Interrupt, "Interrupted")
}
