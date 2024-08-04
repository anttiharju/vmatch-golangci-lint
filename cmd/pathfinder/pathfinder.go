package pathfinder

import (
	"os"
	"os/exec"
	"strings"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exit"
	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
)

func GetGoPath() string {
	goBinPath, err := exec.LookPath("go")
	if err != nil {
		exit.WithMessage(exitcode.NoGo, "Cannot find Go in PATH")
	}

	goPathBytes, err := exec.Command(goBinPath, "env", "GOPATH").Output()
	if err != nil {
		exit.WithMessage(exitcode.GoPathIssue, "Cannot get GOPATH")
	}

	goPath := string(goPathBytes)

	return strings.TrimSpace(goPath)
}

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
