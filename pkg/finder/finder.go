package finder

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/anttiharju/vmatch/pkg/filefinder"
)

func GetVersion(filename string) (string, error) {
	filePath, err := filefinder.Locate(filename)
	if err != nil {
		return "", fmt.Errorf("cannot find version file '%s': %w", filename, err)
	}

	return ReadLinterVersion(filePath)
}

func ReadLinterVersion(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot read version file '%s': %w", filePath, err)
	}

	rawContent := strings.TrimSpace(string(content))

	return validateLinterVersion(rawContent)
}

var versionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func validateLinterVersion(version string) (string, error) {
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid version format '%s'", version)
	}

	return version, nil
}
