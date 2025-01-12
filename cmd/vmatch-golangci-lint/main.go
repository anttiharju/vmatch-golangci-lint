package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anttiharju/vmatch/internal/lintapp"
	"github.com/anttiharju/vmatch/pkg/exit"
	"github.com/anttiharju/vmatch/pkg/exit/exitcode"
)

func main() {
	ctx := context.Background()

	go listenInterrupts()

	lintapp := lintapp.NewApp(".golangci-version")
	exitCode := lintapp.Run(ctx)
	exit.Now(exitCode)
}

func listenInterrupts() {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	<-interruptCh
	exit.WithNewlineMessage(exitcode.Interrupt, "Interrupted")
}
