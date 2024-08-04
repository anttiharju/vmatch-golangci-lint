package pathfinder

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anttiharju/homebrew-golangci-lint-updater/cmd/exitcode"
)

func GetGoPath() string {
	goBinPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("Cannot find Go in PATH")
		os.Exit(exitcode.NoGo)
	}

	goPathBytes, err := exec.Command(goBinPath, "env", "GOPATH").Output()
	if err != nil {
		fmt.Println("Cannot get GOPATH")
		os.Exit(exitcode.GoPathIssue)
	}

	goPath := string(goPathBytes)

	return strings.TrimSpace(goPath)
}

func GetBinPath() string {
	binPath, err := os.Executable()
	if err != nil {
		os.Exit(exitcode.BinPathIssue)
	}

	return binPath
}

func GetBinDir() string {
	binPath := GetBinPath()
	binDir := binPath[:strings.LastIndex(binPath, string(os.PathSeparator))]

	return binDir
}

func GetWorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(exitcode.WorkDirIssue)
	}

	return wd
}
