package emailapp

import (
	sender "diaryhub/sso-service/internal/smtp"
	"log/slog"
)

type App struct {
	log         *slog.Logger
	EmailSender *sender.Sender
}

func New(
	log *slog.Logger,
	smtpAddr string,
	smtpHost string,
	smtpSender string,
	smtpUsername string,
	smtpPassword string,
) *App {

	EmailSender := sender.New(
		smtpAddr,
		smtpHost,
		smtpSender,
		smtpUsername,
		smtpPassword,
	)

	return &App{
		log:         log,
		EmailSender: EmailSender,
	}
}
