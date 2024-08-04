package exit

import (
	"fmt"
	"os"
)

// os.Exit doesn't respect defers so it's kind of neat to wrap it in a package to avoid the footgun

func Now(exitCode int) {
	os.Exit(exitCode)
}

func WithMessage(exitCode int, message string) {
	fmt.Println("golangci-lint-updater: " + message)
	os.Exit(exitCode)
}
