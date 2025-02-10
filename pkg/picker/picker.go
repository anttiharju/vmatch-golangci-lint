package picker

import (
	"context"

	"github.com/anttiharju/vmatch/pkg/finder"
	"github.com/anttiharju/vmatch/pkg/linter"
)

func SelectWrapper(ctx context.Context, args []string) int {
	if len(args) > 0 && args[0] == "go" {
		finder.GetLangVersion()

		return 0
	}

	wrappedLinter := linter.NewWrapper("vmatch-golangci-lint")
	exitCode := wrappedLinter.Run(ctx)

	return exitCode
}
