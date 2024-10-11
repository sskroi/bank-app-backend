package storage

import (
	"bank-app-backend/internal/domain"
	"context"
)

type Storage interface {
	// Authorization
	// ID, PublicID fields will be ignored
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, email, passwordHash string) (domain.User, error)

	// Account
}
