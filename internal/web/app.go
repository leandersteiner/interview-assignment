package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

type App struct {
	mux    *http.ServeMux
	mw     []Middleware
	logger *slog.Logger
}

func NewApp(logger *slog.Logger, mw ...Middleware) *App {
	return &App{
		mux:    http.NewServeMux(),
		mw:     mw,
		logger: logger,
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) Handle(method string, group string, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		v := Values{
			TraceID: "traceid",
			Now:     time.Now(),
		}

		ctx = context.WithValue(ctx, key, &v)

		if err := handler(ctx, w, r); err != nil {
			a.logger.Error("handler", "error", err)
			return
		}
	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + path
	}

	a.mux.HandleFunc(fmt.Sprintf("%s %s", method, finalPath), h)
	s, found := strings.CutSuffix(finalPath, "/")
	if found {
		a.mux.HandleFunc(fmt.Sprintf("%s %s", method, s), h)
	}
}

func (a *App) Get(group string, path string, handler Handler, mw ...Middleware) {
	a.Handle(http.MethodGet, group, path, handler, mw...)
}

func (a *App) Post(group string, path string, handler Handler, mw ...Middleware) {
	a.Handle(http.MethodPost, group, path, handler, mw...)
}
