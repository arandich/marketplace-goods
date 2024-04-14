package config

type PrometheusConfig struct {
	Namespace string `yaml:"namespace" env:"NAMESPACE"`
	Subsystem string `yaml:"subsystem" env:"SUBSYSTEM"`
}
