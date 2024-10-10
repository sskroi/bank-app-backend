package service

import "bank-app-backend/internal/storage"

type Auth interface {
}

type User interface {
}

type Account interface {
}

type Transaction interface {
}

type Service struct {
	Auth
	User
	Account
	Transaction
}

func New(store storage.Storage) *Service {
	return &Service{
		Auth: NewAuthService(store),
	}
}
