package versionfinder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/pathfinder"
)

func GetVersion(filename string) string {
	workDir := pathfinder.GetWorkDir()

	for {
		filePath := filepath.Join(workDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("golangci-lint-updater: Cannot read version file '", filePath, "'")
				os.Exit(exitcode.VersionReadFileIssue)
			}

			return strings.TrimSpace(string(content))
		}

		parentDir := filepath.Dir(workDir)
		if parentDir == workDir {
			break
		}

		workDir = parentDir
	}

	fmt.Println("golangci-lint-updater: Cannot find version file '", filename, "'")
	os.Exit(exitcode.VersionIssue)

	return "" // unreachable but compiler needs it (1.22.5)
}
