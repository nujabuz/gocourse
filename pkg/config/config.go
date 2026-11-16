// Package config
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	ConnectionString       string `yaml:"connectionString"`
	MaxConnections         int    `yaml:"maxConnections"`
	MaxConnectionsIdleTime int    `yaml:"maxConnectionsIdleTime"`
	MaxConnectionsLifetime int    `yaml:"maxConnectionsLifetime"`
	MinIdleConnections     int    `yaml:"minIdleConnections"`
}

func Load(pathFile string) (*Config, error) {
	file, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
