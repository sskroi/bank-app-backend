package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// Authorization
func (self *PgStorage) CreateUser(ctx context.Context, newUser domain.User) error {
	err := self.db.Omit("id", "public_id").WithContext(ctx).Create(&newUser).Error

	var pgErr *pgconn.PgError
	// 23505 == unique_violation error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return domain.ErrUserAlreayExists
	}

	return err
}

func (self *PgStorage) GetUser(ctx context.Context, email, passwordHash string) (domain.User, error) {
	user := domain.User{}

	res := self.db.Where("email = ? AND password_hash = ?", email, passwordHash).WithContext(ctx).Find(&user)

	return user, res.Error
}
