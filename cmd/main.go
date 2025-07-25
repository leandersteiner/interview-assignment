package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/leandersteiner/interview-assignment/internal/calculator"
	"github.com/leandersteiner/interview-assignment/internal/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{}))

	persistFlag := flag.Bool("persist", false, "enable persistence")
	flag.Parse()

	if err := run(logger, *persistFlag); err != nil {
		logger.Error("startup", "error", err)
		os.Exit(1)
	}
}

func run(log *slog.Logger, shouldPersist bool) error {
	var mux http.Handler

	if shouldPersist {
		log.Info("using JSON store")
		store, err := calculator.NewJSONStore()
		if err != nil {
			return fmt.Errorf("failed to create JSON store: %w", err)
		}
		defer func(store *calculator.JSONStore) {
			err := store.Save()
			if err != nil {
				log.Error("could not save store", "error", err)
			}
			log.Info("store saved")
		}(store)
		mux = handlers.NewMux(handlers.MuxConfig{
			Logger: log,
			Store:  store,
		})
	} else {
		log.Info("using in-memory store")
		store := calculator.NewResultStore()
		mux = handlers.NewMux(handlers.MuxConfig{
			Logger: log,
			Store:  store,
		})
	}

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
