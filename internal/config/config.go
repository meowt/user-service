package config

import (
	"project/internal/api"
	"project/internal/storage"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

type Config struct {
	Api     api.Config
	Storage storage.Config
}

func Load(env string) (*Config, error) {
	return &Config{}, nil
}
