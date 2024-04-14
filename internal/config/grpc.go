package config

import "time"

type GrpcConfig struct {
	Network               string        `yaml:"network" env:"NETWORK" default:"tcp"`
	Address               string        `yaml:"address" env:"ADDRESS" default:"0.0.0.0:9091"`
	ConnectionTimeout     time.Duration `yaml:"connection_timeout" env:"CONNECTION_TIMEOUT" default:"5m"`
	MaxConnectionIdle     time.Duration `yaml:"max_connection_idle" env:"MAX_CONNECTION_IDLE" default:"5m"`
	MaxConnectionAge      time.Duration `yaml:"max_connection_age" env:"MAX_CONNECTION_AGE" default:"5m"`
	MaxConnectionAgeGrace time.Duration `yaml:"max_connection_age_grace" env:"MAX_CONNECTION_AGE_GRACE" default:"5m"`
	KeepAliveTimeout      time.Duration `yaml:"keep_alive_timeout" env:"KEEP_ALIVE_TIMEOUT" default:"5m"`
}
