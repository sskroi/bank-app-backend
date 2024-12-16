package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"bank-app-backend/pkg/hasher"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	AccessToken string
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   user.PublicId.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
	})

	accessToken, err := token.SignedString([]byte(s.jwtSignKey))
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{AccessToken: accessToken}, nil
}

func (s *UserService) VerifyAccessToken(ctx context.Context, accessToken string) (uuid.UUID, error) {
	parsedToken, err := jwt.Parse(accessToken,
		func(t *jwt.Token) (interface{}, error) { return []byte(s.jwtSignKey), nil })
	if err != nil {
		return uuid.UUID{}, err
	}

	userPublicIdStr, err := parsedToken.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	userPublicId, err := uuid.Parse(userPublicIdStr)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userPublicId, nil
}

func (s *UserService) Get(ctx context.Context, userPubId uuid.UUID) (domain.User, error) {
	return s.store.GetUser(ctx, userPubId)
}
