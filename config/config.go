package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Env  string `env:"ENV" envDefault:"dev"`
}

func NewConfig() (*Config, error) {
	var c = &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
