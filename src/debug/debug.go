package debug

import (
	"os"

	"github.com/anttiharju/vmatch-golangci-lint/src/finder"
)

// TODO: remove deps to other packages
func getFilePath() string {
	return finder.GetBinDir() + string(os.PathSeparator) + "debug.txt"
}

func WriteToFile(s string) {
	bytes := []byte(s)
	ownerOnlyRW := os.FileMode(0o600)
	_ = os.WriteFile(getFilePath(), bytes, ownerOnlyRW)
}
