package picker

import (
	"github.com/anttiharju/vmatch/pkg/wrapper/language"
	"github.com/anttiharju/vmatch/pkg/wrapper/linter"
)

func firstArgIsGo(args []string) bool {
	return len(args) > 0 && args[0] == "go"
}

func SelectWrapper(args []string) int {
	if firstArgIsGo(args) {
		wrappedLanguage := language.NewWrapper("vmatch-go")
		exitCode := wrappedLanguage.Run(args[1:])

		return exitCode
	}

	wrappedLinter := linter.NewWrapper("vmatch-golangci-lint")
	exitCode := wrappedLinter.Run(args)

	return exitCode
}
