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
	GetUser(ctx context.Context, userPubId uuid.UUID) (domain.User, error)
	// returns domain.ErrInvalidLoginCredentials if there is no user with such email
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	// GetUserId returns domain.ErrUnknownUserPubId if there is no user with such public id
	GetUserId(ctx context.Context, userPubId uuid.UUID) (uint, error)

	// Account
	// ID, Number will be ignored
	CreateAccount(ctx context.Context, account *domain.Account) error
	CloseAccount(ctx context.Context, number uuid.UUID, ownerId uint) error
	GetUserAccounts(ctx context.Context, userId uint, offset, limit int) ([]domain.Account, error)
	GetAccountByNumber(ctx context.Context, number uuid.UUID, ownerId uint, notFoundErr error) (domain.Account, error)

	// Transaction
	CreateTransaction(ctx context.Context,
		senderAcc, receiverAcc domain.Account,
		newTransaction *domain.Transaction) error
	GetAccountTransactions(ctx context.Context,
		accountNumber uuid.UUID,
		offset, limit int) ([]domain.TransactionExtended, error)
}
