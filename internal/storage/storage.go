package storage

import (
	"bank-app-backend/internal/domain"
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	// Authorization
	// ID, PublicID fields will be ignored, returns domain.ErrUserAlreadyExists if email already registered
	CreateUser(ctx context.Context, user domain.User) error
	// returns domain.ErrInvalidLoginCredentials if there is no user with such email
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	// GetUserId returns domain.ErrUnknownUserPubId if there is no user with such public id
	GetUserId(ctx context.Context, userPubId uuid.UUID) (uint, error)

	// Account
	// ID will be ignored
	GetUserAccounts(ctx context.Context, uid uint, offset, limit int) ([]domain.Account, error)
}
