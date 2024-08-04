package config

import "os"

type Config struct {
	InstallDir string
}

func NewConfig() *Config {
	return &Config{InstallDir: string(os.PathSeparator) + "bin"}
}
