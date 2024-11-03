package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
	// "bank-app-backend/pkg/hasher"
	// "time"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/google/uuid"
	// "golang.org/x/net/context"
)

type AccountService struct {
	store storage.Storage
}

func NewAccountService(store storage.Storage) *AccountService {
	return &AccountService{
		store: store,
	}
}

func (s *AccountService) UserAccounts(ctx context.Context, userPubid uuid.UUID,
		offset, limit int) ([]domain.Account, error) {
	var accounts []domain.Account

	userId, err := s.store.GetUserId(ctx, userPubid)
	if err != nil {
		return accounts, err
	}
	fmt.Printf("uid: %d", userId)

	return s.store.GetUserAccounts(ctx, userId, offset, limit)
}
