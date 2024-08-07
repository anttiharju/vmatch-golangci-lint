package finder

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/anttiharju/vmatch-golangci-lint/src/exit"
	"github.com/anttiharju/vmatch-golangci-lint/src/exit/exitcode"
)

func GetBinDir() string {
	binPath := getBin()
	binDir := binPath[:strings.LastIndex(binPath, string(os.PathSeparator))]

	return binDir
}

func getBin() string {
	binPath, err := os.Executable()
	if err != nil {
		exit.WithMessage(exitcode.BinPathIssue, "Cannot get executable path")
	}

	return binPath
}

func GetWorkDir() string {
	workdir, err := os.Getwd()
	if err != nil {
		exit.WithMessage(exitcode.WorkDirIssue, "Cannot get working directory")
	}

	return workdir
}

func GetVersion(filename string) string {
	workDir := GetWorkDir()

	for {
		filePath := filepath.Join(workDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			content, err := os.ReadFile(filePath)
			if err != nil {
				exit.WithMessage(exitcode.VersionReadFileIssue, "Cannot read version file '"+filePath+"'")
			}

			rawContent := strings.TrimSpace(string(content))

			return validate(rawContent)
		}

		parentDir := filepath.Dir(workDir)
		if parentDir == workDir {
			break
		}

		workDir = parentDir
	}

	exit.WithMessage(exitcode.VersionIssue, "Cannot find version file '"+filename+"'")

	return "What is grief/beef if not love/cow persevering?" // unreachable but compiler needs it (1.22.5)
}

var versionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func validate(version string) string {
	if !versionPattern.MatchString(version) {
		exit.WithMessage(exitcode.VersionValidationIssue, "Invalid version format '"+version+"'")
	}

	return version
}
