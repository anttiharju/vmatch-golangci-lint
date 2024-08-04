package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/app"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/config"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
)

func main() {
	ctx := context.Background()

	go listenInterrupts()

	config := config.NewConfig()
	app := app.NewApp(config)
	exitCode := app.Run(ctx)
	os.Exit(exitCode)
}

func listenInterrupts() {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	<-interruptCh
	os.Exit(exitcode.Interrupt)
}
