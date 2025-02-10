package picker

import (
	"context"
	"fmt"

	"github.com/anttiharju/vmatch/pkg/linter"
)

func SelectWrapper(ctx context.Context, args []string) int {
	if len(args) > 0 && args[0] == "go" {
		fmt.Println("go")

		return 0
	}

	wrappedLinter := linter.NewWrapper("vmatch-golangci-lint")
	exitCode := wrappedLinter.Run(ctx)

	return exitCode
}
