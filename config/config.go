package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token     string `envconfig:"TOKEN" required:"true"`
	UserAgent string `envconfig:"USER_AGENT" default:"qi-v0.0.1"` // TODO:gitのタグを使うようにする
}

func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("QI", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
