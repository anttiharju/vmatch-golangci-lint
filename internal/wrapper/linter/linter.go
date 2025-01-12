package linter

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/anttiharju/vmatch/internal/exitcode"
	"github.com/anttiharju/vmatch/internal/wrapper"
	"github.com/anttiharju/vmatch/pkg/versionfinder"
)

type WrappedLinter struct {
	desiredVersion string
	installPath    string
}

func getInstallPath(version string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	ps := string(os.PathSeparator)
	installPath := homeDir + ps + ".vmatch" + ps + "golangci-lint" + ps + "v" + version

	return installPath, nil
}

func NewWrapper() *WrappedLinter {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println("wrapperName" + ": " + err.Error())
		os.Exit(exitcode.WorkDirIssue)
	}

	desiredVersion, err := versionfinder.GetVersion(workDir, ".golangci-version")
	if err != nil {
		fmt.Println("wrapperName" + ": " + err.Error())
		os.Exit(exitcode.VersionIssue)
	}

	installPath, err := getInstallPath(desiredVersion)
	if err != nil {
		fmt.Println("wrapperName" + ": " + err.Error())
		os.Exit(exitcode.InstallPathIssue)
	}

	return &WrappedLinter{
		desiredVersion: desiredVersion,
		installPath:    installPath,
	}
}

func (w *WrappedLinter) Run(ctx context.Context) int {
	if w.noBinary() {
		w.install(ctx)
	}

	args := os.Args[1:]
	//nolint:gosec // I don't think a wrapper can avoid G204.
	linter := exec.Command(w.getGolangCILintPath(), args...)
	linterOutput, _ := linter.Output()

	fmt.Print(string(linterOutput))

	return linter.ProcessState.ExitCode()
}

func (w *WrappedLinter) noBinary() bool {
	_, err := os.Stat(w.getGolangCILintPath())

	return os.IsNotExist(err)
}

func (w *WrappedLinter) install(_ context.Context) {
	//nolint:lll // Official binary install command:
	// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
	curl := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"
	pipe := " | "
	sh := "sh -s -- -b "
	command := curl + pipe + sh + w.installPath + " v" + w.desiredVersion
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		w.ExitWithPrint(exitcode.CmdStartIssue, "failed to start command: "+err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		w.ExitWithPrint(exitcode.CmdStartIssue, "failed to wait for command: "+err.Error())
	}
}

func (w *WrappedLinter) getGolangCILintPath() string {
	return w.installPath + string(os.PathSeparator) + "golangci-lint"
}

// os.Exit() does not respect defer so it's neat to wrap its usage in methods.

func (w *WrappedLinter) Exit(exitCode int) {
	os.Exit(exitCode)
}

func (w *WrappedLinter) ExitWithPrint(exitCode int, message string) {
	fmt.Print("wrapperName" + ": " + message)
	os.Exit(exitCode)
}

func (w *WrappedLinter) ExitWithPrintln(exitCode int, message string) {
	fmt.Println("\n" + "wrapperName" + ": " + message)
	os.Exit(exitCode)
}

var _ wrapper.Interface = (*WrappedLinter)(nil)
