package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token string `envconfig:"TOKEN" required:"true"`
}

func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("QI", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
