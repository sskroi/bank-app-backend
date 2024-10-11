package service

import (
	"bank-app-backend/internal/storage"
	"bank-app-backend/pkg/hasher"

	"golang.org/x/net/context"
)

type Users interface {
	SignUp(ctx context.Context, input UsersSignUpInput) error
}

type Accounts interface {
}

type Transactions interface {
}

type Services struct {
	Users        Users
	Accounts     Accounts
	Transactions Transactions
}

func New(store storage.Storage, passwordHasher hasher.PasswdHasher) *Services {
	return &Services{
		Users: NewUserService(store, passwordHasher),
	}
}

type UsersSignUpInput struct {
	Email      string
	Password   string
	Passport   string
	Name       string
	Surname    string
	Patronymic *string
}
