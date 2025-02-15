package picker

import (
	"context"

	"github.com/anttiharju/vmatch/pkg/lang"
	"github.com/anttiharju/vmatch/pkg/linter"
)

func firstArgIsGo(args []string) bool {
	return len(args) > 0 && args[0] == "go"
}

func SelectWrapper(ctx context.Context, args []string) int {
	if firstArgIsGo(args) {
		wrappedLang := lang.NewWrapper("vmatch-go")
		exitCode := wrappedLang.Run(ctx, args[1:])

		return exitCode
	}

	wrappedLinter := linter.NewWrapper("vmatch-golangci-lint")
	exitCode := wrappedLinter.Run(ctx, args)

	return exitCode
}
