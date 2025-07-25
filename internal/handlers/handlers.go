package handlers

import (
	"context"
	"github.com/leandersteiner/interview-assignment/internal/calculator"
	"github.com/leandersteiner/interview-assignment/internal/handlers/middleware"
	"github.com/leandersteiner/interview-assignment/internal/web"
	"log/slog"
	"net/http"
)

type MuxConfig struct {
	Logger *slog.Logger
	Store  calculator.Store
}

func NewMux(cfg MuxConfig) http.Handler {
	app := web.NewApp(
		cfg.Logger,
		middleware.Panic(cfg.Logger),
		middleware.Log(cfg.Logger),
		middleware.Errors(cfg.Logger),
	)

	app.Get("", "/healthz", func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		return web.Respond(ctx, w, struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}, http.StatusOK)
	})

	calculator.V1Routes(app, calculator.Config{Logger: cfg.Logger, Store: cfg.Store})

	return app
}
