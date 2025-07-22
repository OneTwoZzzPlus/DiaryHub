package restapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"

	authv1 "diaryhub/sso-service/protos/gen/auth"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	log      *slog.Logger
	portGRPC string
	portREST string
	corsProp string
	mux      http.Handler
	cancel   context.CancelFunc
}

func New(log *slog.Logger, portGRPC string, portREST string, corsProp string) *App {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	mux := runtime.NewServeMux()
	muxCORS := cors(mux, corsProp)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		panic(err)
	}

	return &App{
		log:      log,
		portGRPC: portGRPC,
		portREST: portREST,
		corsProp: corsProp,
		mux:      muxCORS,
		cancel:   cancel,
	}
}

func (a *App) Run() {
	const op = "app.rest.Run"
	a.log.Info("Gateway server starting",
		slog.String("rest_port", a.portREST),
		slog.String("grpc_port", a.portGRPC))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", a.portREST), a.mux); err != nil {
		a.log.Error("Stopping gateway server", slog.String("op", op), slog.String("error", err.Error()))
	}
}

func (a *App) Stop() {
	a.cancel()
}

func cors(h http.Handler, corsProp string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin"), corsProp) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string, corsProp string) bool {
	if corsProp == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(corsProp, origin); matched {
		return true
	}
	return false
}
