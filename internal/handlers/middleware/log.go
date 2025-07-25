package middleware

import (
	"context"
	"github.com/leandersteiner/interview-assignment/internal/web"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func Log(logger *slog.Logger) web.Middleware {
	return func(next web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, err := web.GetValues(ctx)
			if err != nil {
				return err
			}

			logger.Info(
				"request started",
				slog.String("trace_id", v.TraceID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("ip", strings.Split(r.RemoteAddr, ":")[0]),
			)

			err = next(ctx, w, r)

			logger.Info(
				"request finished",
				slog.String("trace_id", v.TraceID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("ip", strings.Split(r.RemoteAddr, ":")[0]),
				slog.Duration("time_taken", time.Since(v.Now)),
				slog.Int("status", v.StatusCode),
			)

			return err
		}
	}
}
