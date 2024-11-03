package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

// Authorization
func (self *PgStorage) CreateUser(ctx context.Context, newUser domain.User) error {
	err := self.db.Omit("id", "public_id").WithContext(ctx).Create(&newUser).Error

	var pgErr *pgconn.PgError
	// 23505 == unique_violation error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return domain.ErrUserAlreadyExists
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

func (self *PgStorage) GetUserId(ctx context.Context, userPubId uuid.UUID) (uint, error) {
	user := domain.User{}
	
	res := self.db.Where("public_id = ?", userPubId).WithContext(
		ctx).Select("id").Find(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, domain.ErrUnknownUserPubId
	}

	return user.ID, nil
}

// func (self *PgStorage) GetUserById(ctx context.Context, uid uint) (domain.User, error) {
// 	user := domain.User{}

// 	res := self.db.Where("id = ?", uid).WithContext(ctx).Find(&user)
// 	if res.Error != nil {
// 		return user, res.Error
// 	}

// 	if res.RowsAffected == 0 {
// 		return user, domain.ErrUnknownUserId
// 	}

// 	return user, nil
// }
