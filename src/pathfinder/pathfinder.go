package pathfinder

import (
	"os"
	"strings"

	"github.com/anttiharju/vmatch-golangci-lint/src/exit"
	"github.com/anttiharju/vmatch-golangci-lint/src/exit/exitcode"
)

func GetBinPath() string {
	binPath, err := os.Executable()
	if err != nil {
		exit.WithMessage(exitcode.BinPathIssue, "Cannot get executable path")
	}

	return binPath
}

func GetBinDir() string {
	binPath := GetBinPath()
	binDir := binPath[:strings.LastIndex(binPath, string(os.PathSeparator))]

	return binDir
}

func GetWorkDir() string {
	workdir, err := os.Getwd()
	if err != nil {
		exit.WithMessage(exitcode.WorkDirIssue, "Cannot get working directory")
	}

	return workdir
}
