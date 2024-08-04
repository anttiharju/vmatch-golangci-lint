package config

type Config struct {
	InstallDir string
}

func NewConfig() *Config {
	return &Config{InstallDir: "/bin"}
}
