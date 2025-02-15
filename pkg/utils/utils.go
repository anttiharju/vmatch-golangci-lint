package utils

import (
	"fmt"
	"os"
)

func GenerateInstallPath(name, version string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get install path: %w", err)
	}

	ps := string(os.PathSeparator)
	installPath := homeDir + ps + ".vmatch" + ps + name + ps + "v" + version

	return installPath, nil
}
