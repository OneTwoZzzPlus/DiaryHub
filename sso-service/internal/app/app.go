package app

import (
	emailapp "diaryhub/sso-service/internal/app/email"
	grpcapp "diaryhub/sso-service/internal/app/grpc"
	restapp "diaryhub/sso-service/internal/app/rest"
	storageapp "diaryhub/sso-service/internal/app/storage"
	authservice "diaryhub/sso-service/internal/services/auth"
	"log/slog"
	"time"
)

type App struct {
	GRPCApp     *grpcapp.App
	RESTApp     *restapp.App
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

	RESTApp := restapp.New(log, "9090", "7070", "*")

	GRPCApp := grpcapp.New(log, grpcPort, AuthService)

	return &App{
		GRPCApp:     GRPCApp,
		RESTApp:     RESTApp,
		AuthService: AuthService,
		StorageApp:  StorageApp,
		EmailApp:    EmailApp,
	}
}

func (a *App) Stop() {
	a.GRPCApp.Stop()
	a.RESTApp.Stop()
	a.StorageApp.Disconnect()
}
