package config

type LoggerConfig struct {
	Level string `yaml:"level" env:"LEVEL" default:"info"`
}
