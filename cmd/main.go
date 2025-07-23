package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{}))

	if err := run(logger); err != nil {
		logger.Error("startup", "error", err)
		os.Exit(1)
	}
}

func run(log *slog.Logger) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	server := &http.Server{
		Addr:        "127.0.0.1:8080",
		Handler:     http.TimeoutHandler(mux, 5*time.Second, "request timed out"),
		IdleTimeout: 30 * time.Second,
	}

	serverError := make(chan error, 1)

	go func() {
		log.Info("starting server", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverError <- fmt.Errorf("error starting server: %w", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		return err
	case <-stop:
		log.Info("graceful shutdown initiated")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("error gracefully shutting down server: %w", err)
		}
		log.Info("graceful shutdown complete")
	}

	return nil
}
