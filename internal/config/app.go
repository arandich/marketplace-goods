package config

type AppConfig struct {
	Name   string `yaml:"name" env:"NAME" default:"app"`
	Source string `yaml:"source" env:"SOURCE" required:"true"`
}
