package exit

import (
	"fmt"
	"os"
)

const prefix = "vmatch-golangci-lint: "

// os.Exit doesn't respect defers so it's kind of neat to wrap it in a package to avoid the footgun

func Now(exitCode int) {
	os.Exit(exitCode)
}

func WithNewlineMessage(exitCode int, message string) {
	fmt.Println("\n" + prefix + message)
	os.Exit(exitCode)
}

func WithMessage(exitCode int, message string) {
	fmt.Print(prefix + message)
	os.Exit(exitCode)
}
