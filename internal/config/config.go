package config

import (
	"os"

	"github.com/BurntSushi/toml"

	"bank-app-backend/internal/server"
	"bank-app-backend/internal/storage/postgres"
)

const (
	defaultConfigPath = "configs/config.toml"
)

type (
	Config struct {
		Postgres postgres.Config `toml:"postgres"`
		Server   server.Config   `toml:"server"`
		Auth     Auth            `toml:"auth"`
	}

	Auth struct {
		JwtSignKey string `toml:"jwtsignkey"`
	}
)

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("BANK_APP_CONFIG_PATH")
	if configPath == "" {
		configPath = defaultConfigPath
	}

	cfg := new(Config)

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
