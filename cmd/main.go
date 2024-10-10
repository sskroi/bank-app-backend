package main

import (
	"bank-app-backend/internal/config"
	"bank-app-backend/internal/handler"
	"bank-app-backend/internal/server"
	"bank-app-backend/internal/service"
	"bank-app-backend/internal/storage/postgres"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("can't load config: %s", err)
	}

	store, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("can't connect to database: %s", err)
	}

	service := service.New(store)
	handler := handler.New(service)

	server := new(server.Server)
	go func() {
		err := server.Run(cfg.Server, handler.InitRoutes())
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err)
		}
	}()

	log.Printf("app started on port: %s", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("app shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := store.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}
