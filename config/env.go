package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	MySqlHost            string `envconfig:"MYSQL_HOST"`
	MySqlPort            string `envconfig:"MYSQL_PORT"`
	MySqlDB              string `envconfig:"MYSQL_DB"`
	MySqlUser            string `envconfig:"MYSQL_USER"`
	MySqlPass            string `envconfig:"MYSQL_PASSWORD"`
	MySQLConnMaxLifetime int    `envconfig:"MYSQL_CONN_MAX_LIFETIME" default:"55" required:"true"`
}

func LoadEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	return &env, nil
}
