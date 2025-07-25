package calculator

import (
	"github.com/leandersteiner/interview-assignment/internal/web"
	"log/slog"
)

type Config struct {
	Logger *slog.Logger
	Store  Store
}

func V1Routes(app *web.App, cfg Config) {
	const version = "api/v1/calculator"

	service := NewService(4, cfg.Store)
	handler := NewHandler(service, cfg.Store)

	app.Post(version, "/addition", handler.Addition)
	app.Post(version, "/subtraction", handler.Subtraction)
	app.Post(version, "/multiplication", handler.Multiplication)
	app.Post(version, "/division", handler.Division)
	app.Get(version, "/recent", handler.GetRecent)
}
