package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"bank-app-backend/pkg/hasher"

	"golang.org/x/net/context"
)

type UserService struct {
	store      storage.Storage
	pswdHasher hasher.PasswdHasher
}

func NewUserService(store storage.Storage, passwordHasher hasher.PasswdHasher) *UserService {
	return &UserService{
		store:      store,
		pswdHasher: passwordHasher,
	}
}

func (self *UserService) SignUp(ctx context.Context, input UsersSignUpInput) error {
	hashedPassword, err := self.pswdHasher.Hash(input.Password)
	if err != nil {
		return err
	}

	if err := self.store.CreateUser(ctx, domain.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Passport:     input.Passport,
		Name:         input.Name,
		Surname:      input.Surname,
		Patronymic:   input.Patronymic,
	}); err != nil {
		return err
	}

	// maybe send verification email

	return nil
}
