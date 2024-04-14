package config

import "time"

type HttpConfig struct {
	Network           string        `yaml:"network" env:"NETWORK" default:"tcp"`
	Address           string        `yaml:"address" env:"ADDRESS" default:":8080"`
	ReadTimeout       time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" default:"5m"`
	WriteTimeout      time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" default:"5m"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" default:"5m"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT" default:"5m"`
	ProfilingEnabled  bool          `yaml:"profiling_enabled" env:"PROFILING_ENABLED" default:"true"`
}
