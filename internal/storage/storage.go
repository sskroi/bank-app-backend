package storage

import "bank-app-backend/internal/domain"


type Storage interface {
	// Authorization
	CreateUser(email, passwordHash, passport, name, surname string, patronymic *string) (domain.User, error)
	GetUser(email, passwordHash string) (domain.User, error)

	// Account
}

