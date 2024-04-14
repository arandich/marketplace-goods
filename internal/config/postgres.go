package config

type PostgresConfig struct {
	Host     string `yaml:"host" env:"HOST" required:"true"`
	Port     string `yaml:"port" env:"PORT" required:"true"`
	User     string `yaml:"user" env:"USER" required:"true"`
	Password string `yaml:"password" env:"PASSWORD" required:"true"`
	Database string `yaml:"database" env:"DATABASE" default:"db"`
}
