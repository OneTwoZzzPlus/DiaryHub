package main

import (
	"diaryhub/auth-service/internal/config"
	"diaryhub/auth-service/internal/storage/postgresql"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Starting auth-service", slog.String("env", cfg.Env))
	log.Debug("Debug messages enabled")
	log.Debug(fmt.Sprintf("Config: %s", cfg))

	_, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage", slog.String("error", err.Error()))
		return
	}
	log.Info("Storage connected")

	// TODO: init router
	// TODO: run server
}

const (
	envLocal       = "local"
	envDevelopment = "dev"
	envProduction  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDevelopment:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduction:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
