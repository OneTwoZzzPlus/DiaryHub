package app

import (
	emailapp "diaryhub/sso-service/internal/app/email"
	grpcapp "diaryhub/sso-service/internal/app/grpc"
	storageapp "diaryhub/sso-service/internal/app/storage"
	authservice "diaryhub/sso-service/internal/services/auth"
	"log/slog"
	"time"
)

type App struct {
	GRPCApp     *grpcapp.App
	AuthService *authservice.Auth
	StorageApp  *storageapp.App
	EmailApp    *emailapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
	smtpAddr string,
	smtpHost string,
	smtpSender string,
	smtpUsername string,
	smtpPassword string,
) *App {

	StorageApp := storageapp.MustConnect(log, storagePath)

	EmailApp := emailapp.New(
		log,
		smtpAddr,
		smtpHost,
		smtpSender,
		smtpUsername,
		smtpPassword,
	)

	AuthService := authservice.New(
		log,
		tokenTTL,
		StorageApp.Storage,
		StorageApp.Storage,
		StorageApp.Storage,
		EmailApp.EmailSender,
	)

	GRPCApp := grpcapp.New(log, grpcPort, AuthService)

	return &App{GRPCApp: GRPCApp, AuthService: AuthService, StorageApp: StorageApp, EmailApp: EmailApp}
}

func (a *App) Stop() {
	a.GRPCApp.Stop()
	a.StorageApp.Disconnect()
}
