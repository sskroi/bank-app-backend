package storage

import (
	"bank-app-backend/internal/domain"
	"context"
)

type Storage interface {
	// Authorization
	// ID, PublicID fields will be ignored, returns domain.ErrUserAlreayExists if email already registered
	CreateUser(ctx context.Context, user domain.User) error
	// returns domain.ErrInvalidLoginCredentials if there is no user with such email
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)

	// Account
}
