package main

import (
	"diaryhub/auth-service/internal/app"
	"diaryhub/auth-service/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	blog "github.com/charmbracelet/log"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("Starting auth-service", slog.String("env", cfg.Env))
	log.Debug("Debug messages enabled")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.StorageApp.MustRun()
	go application.GRPCApp.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("STOPPING auth-service", slog.Any("signal", sign))

	application.GRPCApp.Stop()
	application.StorageApp.Stop()
	log.Info("auth-service stopped")
}

const (
	envLocal       = "local"
	envDevelopment = "dev"
	envProduction  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(blog.New(os.Stdout))
	case envDevelopment:
		logger = slog.New(blog.New(os.Stdout))
	case envProduction:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
