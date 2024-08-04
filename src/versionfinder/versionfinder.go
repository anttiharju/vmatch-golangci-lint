package versionfinder

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/anttiharju/homebrew-golangci-lint-updater/src/exit"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/exit/exitcode"
	"github.com/anttiharju/homebrew-golangci-lint-updater/src/pathfinder"
)

func GetVersion(filename string) string {
	workDir := pathfinder.GetWorkDir()

	for {
		filePath := filepath.Join(workDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			content, err := os.ReadFile(filePath)
			if err != nil {
				exit.WithMessage(exitcode.VersionReadFileIssue, "Cannot read version file '"+filePath+"'")
			}

			return strings.TrimSpace(string(content))
		}

		parentDir := filepath.Dir(workDir)
		if parentDir == workDir {
			break
		}

		workDir = parentDir
	}

	exit.WithMessage(exitcode.VersionIssue, "Cannot find version file '"+filename+"'")

	return "" // unreachable but compiler needs it (1.22.5)
}
