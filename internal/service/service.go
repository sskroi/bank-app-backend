package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"bank-app-backend/pkg/hasher"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Users interface {
	SignUp(ctx context.Context, input UsersSignUpInput) error
	SignIn(ctx context.Context, email, password string) (Tokens, error)
	// VerifyAccessToken verifies token and return user's public id if token is valid
	VerifyAccessToken(ctx context.Context, accessToken string) (uuid.UUID, error)
}

type Accounts interface {
	// Create(ctx context.Context, uid int) error
	UserAccounts(ctx context.Context, userPubId uuid.UUID, offset, limit int) ([]domain.Account, error)
	// GetBalance(ctxt context.Context) (decimal.Decimal, error)
}

type Transactions interface {
}

type Services struct {
	Users        Users
	Accounts     Accounts
	Transactions Transactions
}

func New(store storage.Storage, passwordHasher hasher.PasswdHasher, jwtSignKey string) *Services {
	return &Services{
		Users: NewUserService(store, passwordHasher, jwtSignKey),
		Accounts: NewAccountService(store),
	}
}
