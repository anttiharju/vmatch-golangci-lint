package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anttiharju/homebrew-golangci-lint-updater/src/app"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/exit"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/exit/exitcode"
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
	exit.WithNewlineMessage(exitcode.Interrupt, "Interrupted")
}
