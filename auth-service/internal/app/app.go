package app

import (
	grpcapp "diaryhub/auth-service/internal/app/grpc"
	storageapp "diaryhub/auth-service/internal/app/storage"
	"log/slog"
	"time"
)

type App struct {
	GRPCApp    *grpcapp.App
	StorageApp *storageapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	StorageApp := storageapp.New(log, storagePath)

	GRPCApp := grpcapp.New(log, grpcPort)

	return &App{GRPCApp: GRPCApp, StorageApp: StorageApp}
}
