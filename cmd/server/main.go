package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/musician-production-suite/internal/api"
	"github.com/example/musician-production-suite/internal/config"
	"github.com/example/musician-production-suite/internal/jobs"
	"github.com/example/musician-production-suite/internal/utils"
	"github.com/example/musician-production-suite/pkg/audio"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-healthcheck" {
		if err := healthcheck(); err != nil {
			os.Exit(1)
		}
		return
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	cfg, err := config.Load()
	if err != nil {
		utils.HandleErrorOrLogWithMessages(logger, err, "load config", "")
		os.Exit(1)
	}

	pipeline := audio.NewPipeline(logger)
	store, err := jobs.NewFileStore(cfg.StorageDir)
	if err != nil {
		utils.HandleErrorOrLogWithMessages(logger, err, "create job store", "")
		os.Exit(1)
	}
	runner := jobs.NewRunner(store, pipeline, logger)
	server := &http.Server{
		Addr:              cfg.APIAddr,
		Handler:           api.NewRouter(api.Dependencies{Config: cfg, Store: store, Runner: runner, Version: version, Logger: logger}),
		ReadHeaderTimeout: 10 * time.Second,
	}

	errs := make(chan error, 1)
	go func() {
		logger.Info("server_starting", "addr", cfg.APIAddr, "version", version)
		errs <- server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		if err != nil && err != http.ErrServerClosed {
			utils.HandleErrorOrLogWithMessages(logger, err, "server failed", "")
			os.Exit(1)
		}
	case sig := <-signals:
		logger.Info("shutdown_signal", "signal", sig.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		utils.HandleErrorOrLogWithMessages(logger, err, "graceful shutdown", "")
		os.Exit(1)
	}
	logger.Info("server_stopped")
}

func healthcheck() error {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 2*time.Second)
	if err != nil {
		return err
	}
	return conn.Close()
}
