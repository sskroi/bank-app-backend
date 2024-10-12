package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"bank-app-backend/pkg/hasher"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

type UserService struct {
	store      storage.Storage
	pswdHasher hasher.PasswdHasher
	jwtSignKey string
}

func NewUserService(store storage.Storage, passwordHasher hasher.PasswdHasher, jwtSignKey string) *UserService {
	return &UserService{
		store:      store,
		pswdHasher: passwordHasher,
		jwtSignKey: jwtSignKey,
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

type Tokens struct {
	AccessToken  string
	// RefreshToken string
}

func (s *UserService) SignUp(ctx context.Context, input UsersSignUpInput) error {
	hashedPassword, err := s.pswdHasher.Hash(input.Password)
	if err != nil {
		return err
	}

	if err := s.store.CreateUser(ctx, domain.User{
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

func (s *UserService) SignIn(ctx context.Context, email, password string) (Tokens, error) {
	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return Tokens{}, err
	}

	if !s.pswdHasher.Check(user.PasswordHash, password) {
		return Tokens{}, domain.ErrInvalidLoginCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.PublicId,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	})

	accessToken, err := token.SignedString([]byte(s.jwtSignKey))
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{AccessToken: accessToken}, nil
}
