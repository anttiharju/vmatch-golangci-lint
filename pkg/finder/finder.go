package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetLangVersion(filename string) (string, error) {
	filePath, err := locateFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot find lang version file '%s': %w", filename, err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read version file '%s': %w", filePath, err)
	}

	return readLangVersion(content, filePath)
}

func GetLinterVersion(filename string) (string, error) {
	filePath, err := locateFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot find linter version file '%s': %w", filename, err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read version file '%s': %w", filePath, err)
	}

	return readLinterVersion(content, filePath)
}

func locateFile(filename string) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current working directory: %w", err)
	}

	for {
		filePath := filepath.Join(workDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		}

		parentDir := filepath.Dir(workDir)
		if parentDir == workDir {
			break
		}

		workDir = parentDir
	}

	return "", fmt.Errorf("cannot find version file '%s'", filename)
}

func readLangVersion(content []byte, filePath string) (string, error) {
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			trimmed := strings.TrimPrefix(line, "go ")

			return validateVersion(trimmed)
		}
	}

	return "", fmt.Errorf("cannot find go version in file '%s'", filePath)
}

func readLinterVersion(content []byte, _ string) (string, error) {
	trimmed := strings.TrimSpace(string(content))

	return validateVersion(trimmed)
}

var versionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func validateVersion(version string) (string, error) {
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid version format '%s'", version)
	}

	return version, nil
}
