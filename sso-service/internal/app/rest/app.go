package restapp

import (
	"context"
	"log/slog"
	"net/http"

	authv1 "diaryhub/sso-service/protos/gen/auth"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	log      *slog.Logger
	portGRPC string
	portREST string
	mux      *runtime.ServeMux
	cancel   context.CancelFunc
}

func New(log *slog.Logger, portGRPC string, portREST string) *App {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		panic(err)
	}

	return &App{
		log:      log,
		portGRPC: portGRPC,
		portREST: portREST,
		mux:      mux,
		cancel:   cancel,
	}
}

func (a *App) Run() {
	a.log.Info("Gateway server listening at 7070")
	if err := http.ListenAndServe(":7070", a.mux); err != nil {
		a.log.Error("Stopping gateway server", slog.String("error", err.Error()))
	}
}

func (a *App) Stop() {
	a.cancel()
}
