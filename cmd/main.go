package main

import (
	"bank-app-backend/internal/config"
	"bank-app-backend/internal/handler"
	"bank-app-backend/internal/server"
	"bank-app-backend/internal/service"
	"bank-app-backend/internal/storage/postgres"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title			Backend part of educational banking application
// @version		1.0
// @host			localhost:8080
// @BasePath		/api/v1
// @license.name	MIT
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("can't load config", "err", err)
		return
	}

	store, err := postgres.New(cfg.Postgres)
	if err != nil {
		slog.Error("can't connect to database", "err", err)
		return
	}

	service := service.New(store)
	handler := handler.New(service)

	server := new(server.Server)
	go func() {
		err := server.Run(cfg.Server, handler.InitRoutes())
		if err != nil && err != http.ErrServerClosed {
			slog.Error("error occured while running http server", "err", err)
		}
	}()

	slog.Info("app started", "port", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("app shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down", "err", err)
	}

	if err := store.Close(); err != nil {
		slog.Error("error occured on db connection close", "err", err)
	}
}
