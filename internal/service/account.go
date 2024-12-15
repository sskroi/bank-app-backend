package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"context"

	"github.com/google/uuid"
)

type AccountService struct {
	store storage.Storage
}

func NewAccountService(store storage.Storage) *AccountService {
	return &AccountService{
		store: store,
	}
}

func (s *AccountService) Create(
		ctx context.Context, userPubId uuid.UUID, currency string) (uuid.UUID, error) {
	var accountNumber uuid.UUID
	
	userId, err := s.store.GetUserId(ctx, userPubId)
	if err != nil {
		return accountNumber, err
	}

	newAccount := domain.Account{
		OwnerId: userId,
		Currency: currency,
	}

	if err := s.store.CreateAccount(ctx, &newAccount); err != nil {
		return accountNumber, err
	}

	return newAccount.Number, err
}

func (s *AccountService) Close(ctx context.Context,
		userPubId, number uuid.UUID) error {
	userId, err := s.store.GetUserId(ctx, userPubId)
	if err != nil {
		return err
	}
	return s.store.CloseAccount(ctx, number, userId)
}

func (s *AccountService) UserAccounts(ctx context.Context, userPubId uuid.UUID,
		offset, limit int) ([]domain.Account, error) {
	var accounts []domain.Account

	userId, err := s.store.GetUserId(ctx, userPubId)
	if err != nil {
		return accounts, err
	}

	return s.store.GetUserAccounts(ctx, userId, offset, limit)
}
