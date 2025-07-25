package middleware

import (
	"context"
	"github.com/leandersteiner/interview-assignment/internal/web"
	"log/slog"
	"net/http"
)

func Panic(logger *slog.Logger) web.Middleware {
	return func(next web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			var pErr error
			defer func() {
				v, err := web.GetValues(ctx)
				if err != nil {
					pErr = err
					return
				}
				if err := recover(); err != nil {
					logger.Error("panic", "trace_id", v.TraceID, "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr, "error", err)
					pErr = web.NewError(http.StatusInternalServerError, "There was an internal server error")
				}
			}()
			if pErr != nil {
				return pErr
			}
			return next(ctx, w, r)
		}
	}
}
