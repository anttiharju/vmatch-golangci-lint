package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch-golangci-lint/src/config"
	"github.com/anttiharju/vmatch-golangci-lint/src/pathfinder"
	"github.com/anttiharju/vmatch-golangci-lint/src/versionfinder"
)

type App struct {
	config         *config.Config
	goPath         string
	desiredVersion string
}

func NewApp(config *config.Config) *App {
	return &App{
		config:         config,
		goPath:         pathfinder.GetGoPath(),
		desiredVersion: versionfinder.GetVersion(config.VersionFileName),
	}
}

func (a *App) Run(ctx context.Context) int {
	a.install(ctx)
	args := os.Args[1:]
	linter := exec.Command(a.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (a *App) install(_ context.Context) {
	//nolint:lll // Official binary install command, it is what it is
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	ps := string(os.PathSeparator)
	installPath := pathfinder.GetBinDir() + ps + a.desiredVersion + ps + " " + a.desiredVersion // TODO: do v/1.59.1/ instead of v1.59.1
	cmd := curl + pipe + sh + installPath
	execCmd := exec.Command("sh", "-c", cmd)
	output, _ := execCmd.Output()
	fmt.Println(string(output))
}

func (a *App) getGolangCILintPath() string {
	return a.goPath + a.config.InstallDir + string(os.PathSeparator) + "golangci-lint"
}
