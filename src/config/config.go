package config

import "os"

type Config struct {
	InstallDir      string
	VersionFileName string
}

func NewConfig() *Config {
	return &Config{
		InstallDir:      string(os.PathSeparator) + "bin", // TODO: switch to a subdir under vmatch-golangci-lint bin location
		VersionFileName: ".golangci-version",
	}
}
