package middleware

import (
	"context"
	"errors"
	"github.com/leandersteiner/interview-assignment/internal/web"
	"log/slog"
	"net/http"
)

func Errors(logger *slog.Logger) web.Middleware {
	return func(handler web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, err := web.GetValues(ctx)
			if err != nil {
				return err
			}

			err = handler(ctx, w, r)
			if err != nil {
				logger.Error("request error", "trace_id", v.TraceID, "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr, "error", err)

				var webErr *web.Error
				if errors.As(err, &webErr) {
					err = web.Respond(ctx, w, webErr, webErr.Status)
					return nil
				}

				err = web.Respond(ctx, w, err, http.StatusInternalServerError)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
}
