package picker

import (
	"context"
	"fmt"
	"runtime"

	"github.com/anttiharju/vmatch/pkg/filefinder"
	"github.com/anttiharju/vmatch/pkg/linter"
)

func SelectWrapper(ctx context.Context, args []string) int {
	if len(args) > 0 && args[0] == "go" {
		filePath, _ := filefinder.Locate("go.mod")
		fmt.Println("Found go.mod at", filePath)

		version := "1.23.5"
		fmt.Println("https://go.dev/dl/go" + version + "." + runtime.GOOS + "-" + runtime.GOARCH + ".tar.gz")

		return 0
	}

	wrappedLinter := linter.NewWrapper("vmatch-golangci-lint")
	exitCode := wrappedLinter.Run(ctx)

	return exitCode
}
