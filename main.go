package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/anttiharju/vmatch/pkg/exitcode"
	"github.com/anttiharju/vmatch/pkg/picker"
)

func main() {
	go interruptListener(os.Interrupt)

	ctx := context.Background()
	exitCode := picker.SelectWrapper(ctx, os.Args[1:])

	os.Exit(exitCode)
}

func interruptListener(signals ...os.Signal) {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, signals...)
	<-interruptCh
	fmt.Println("\nvmatch: interrupted") // leading \n to have ^C appear on its own line
	os.Exit(exitcode.Interrupt)
}
