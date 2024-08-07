package config

type Config struct {
	VersionFileName string
}

func NewConfig() *Config {
	return &Config{
		VersionFileName: ".golangci-version",
	}
}
