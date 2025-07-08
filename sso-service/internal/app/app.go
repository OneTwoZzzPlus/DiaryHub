package app

import (
	grpcapp "diaryhub/sso-service/internal/app/grpc"
	storageapp "diaryhub/sso-service/internal/app/storage"
	authservice "diaryhub/sso-service/internal/services/auth"
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

	StorageApp := storageapp.MustConnect(log, storagePath)

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

func (a *App) Stop() {
	a.GRPCApp.Stop()
	a.StorageApp.Disconnect()
}
