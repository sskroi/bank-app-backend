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

func (self *PgStorage) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}

	res := self.db.Where("email = ?", email).WithContext(ctx).Find(&user)
	if res.Error != nil {
		return user, res.Error
	}

	if res.RowsAffected == 0 {
		return user, domain.ErrInvalidLoginCredentials
	}

	return user, nil
}
