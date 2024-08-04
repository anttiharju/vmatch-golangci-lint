package config

import "os"

type Config struct {
	InstallDir      string
	VersionFileName string
}

func NewConfig() *Config {
	return &Config{
		InstallDir:      string(os.PathSeparator) + "bin", // TODO: switch bin to golangci-lint-updater
		VersionFileName: ".golangci-version",
	}
}
