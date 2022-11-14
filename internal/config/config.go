package config

import "github.com/kelseyhightower/envconfig"

// Config stores service configuration value
type Config struct {
	HTTPServerPort int `envconfig:"HTTP_SERVER_PORT" default:"8080"`
}

// New initializes new config by loading the value from environment variables.
func New() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
