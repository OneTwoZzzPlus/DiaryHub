package storageapp

import (
	"diaryhub/auth-service/internal/storage/postgresql"
	"fmt"

	"log/slog"
)

type App struct {
	log         *slog.Logger
	storagePath string
	Storage     *postgresql.Storage
}

func New(log *slog.Logger, storagePath string) *App {
	return &App{log: log, storagePath: storagePath}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.storage.Connect"
	log := a.log.With(slog.String("op", op))

	log.Info("Connecting to storage...")
	strg, err := postgresql.New(a.storagePath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Storage connected")
	a.Storage = strg
	return nil
}

func (a *App) Stop() {
	const op = "app.storage.Stop"
	log := a.log.With(slog.String("op", op))

	log.Info("STOPPING storage")
	a.Storage.Close()
}
