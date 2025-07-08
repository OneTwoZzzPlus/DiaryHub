package app

import (
	grpcapp "diaryhub/auth-service/internal/app/grpc"
	storageapp "diaryhub/auth-service/internal/app/storage"
	authservice "diaryhub/auth-service/internal/services/auth"
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

	AuthService := authservice.New(
		log,
		StorageApp.Storage,
		StorageApp.Storage,
		StorageApp.Storage,
		tokenTTL,
	)

	GRPCApp := grpcapp.New(log, grpcPort, AuthService)

	return &App{GRPCApp: GRPCApp, StorageApp: StorageApp}
}
