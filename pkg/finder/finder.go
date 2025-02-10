package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

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

func GetLangVersion() (string, error) {
	const filename = "go.mod"

	filePath, err := locateFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot find lang version file '%s': %w", filename, err)
	}

	version, _ := readLangVersion(filePath)

	fmt.Println("https://go.dev/dl/go" + version + "." + runtime.GOOS + "-" + runtime.GOARCH + ".tar.gz")

	return version, nil
}

func readLangVersion(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read version file '%s': %w", filePath, err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			rawVersion := strings.TrimPrefix(line, "go ")

			return validateLangVersion(rawVersion)
		}
	}

	return "", fmt.Errorf("cannot find go version in file '%s'", filePath)
}

func GetLinterVersion() (string, error) {
	const filename = ".golangci-version"

	filePath, err := locateFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot find linter version file '%s': %w", filename, err)
	}

	return readLinterVersion(filePath)
}

func readLinterVersion(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read version file '%s': %w", filePath, err)
	}

	rawContent := strings.TrimSpace(string(content))

	return validateLinterVersion(rawContent)
}

var versionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func validateLangVersion(version string) (string, error) {
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid version format '%s'", version)
	}

	return version, nil
}

func validateLinterVersion(version string) (string, error) {
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid version format '%s'", version)
	}

	return version, nil
}
