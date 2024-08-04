package config

import "os"

type Config struct {
	InstallDir      string
	VersionFileName string
}

func NewConfig() *Config {
	cfg := &Config{
		InstallDir:      string(os.PathSeparator) + "bin", // TODO: switch bin to golangci-lint-updater
		VersionFileName: ".golangci-version",
	}

	return cfg
}
