package config

import "os"

type Config struct {
	InstallDir      string
	VersionFileName string
}

func NewConfig() *Config {
	cfg := &Config{
		InstallDir:      string(os.PathSeparator) + "bin",
		VersionFileName: ".golangci-version",
	}

	return cfg
}
