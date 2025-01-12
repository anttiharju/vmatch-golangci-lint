package pathfinder

import (
	"fmt"
	"os"
	"strings"
)

func GetUserHomeDirPath(version string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	ps := string(os.PathSeparator)
	installPath := homeDir + ps + ".vmatch" + ps + "golangci-lint" + ps + "v" + version

	return installPath, nil
}

func GetExecutablePath() (string, error) {
	binPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	binDir := binPath[:strings.LastIndex(binPath, string(os.PathSeparator))]

	return binDir, nil
}
