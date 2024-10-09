package config

import (
	"github.com/BurntSushi/toml"

	"bank-app-backend/internal/storage/postgres"
)

const (
	configPath = "./configs/config.toml"
)

type Config struct {
	Postgres postgres.Config `toml:"postgres"`
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
