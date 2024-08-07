package debug

import (
	"os"

	"github.com/anttiharju/vmatch-golangci-lint/src/finder"
)

func getFilePath() string {
	return finder.GetBinDir() + string(os.PathSeparator) + "debug.txt"
}

func WriteToFile(s string) {
	d1 := []byte(s)
	_ = os.WriteFile(getFilePath(), d1, 0o600)
}
