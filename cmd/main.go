package main

import (
	"log/slog"
	"os"

	"project/cmd/modules"
	"project/internal/config"
	"project/internal/logger"
	"project/internal/storage"
	"project/pkg/models"
)

var env string

func init() {
	if env = os.Getenv("ENV"); env == "" {
		env = config.EnvDev
	}
}

func main() {
	err := logger.Setup()
	if err != nil {
		slog.Error("Logger pkg", "status", models.StatusError, "err", err)
		return
	}
	slog.Info("Environment", "value", env)
	slog.Info("Config pkg", "status", models.StatusLoaded)

	cfg, err := config.Load(env)
	if err != nil {
		slog.Error("Config pkg", "status", models.StatusError, "err", err)
		return
	}
	slog.Info("Logger pkg", "status", models.StatusLoaded)

	s, err := storage.New(cfg)
	if err != nil {
		slog.Error("Storage pkg", "status", models.StatusError, "err", err)
		return
	}
	slog.Info("Storage pkg", "status", models.StatusLoaded)

	api := modules.Setup(s)
	if err = api.Serve(); err != nil {
		slog.Error("Api serving", "err", err)
		return
	}
}
