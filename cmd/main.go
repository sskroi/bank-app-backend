package main

import (
	"bank-app-backend/internal/config"
	"bank-app-backend/internal/storage/postgres"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("can't load config: %s", err)
	}

	store, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatalf("can't connect to database: %s", err)
	}

	_ = store
}
