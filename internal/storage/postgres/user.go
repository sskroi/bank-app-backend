package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
)

// Authorization
func (self *PgStorage) CreateUser(ctx context.Context, newUser domain.User) (domain.User, error) {
	err := self.db.Omit("id", "public_id").WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return newUser, err
	}

	err = self.db.First(&newUser, newUser.ID).Error
	return newUser, err
}

func (self *PgStorage) GetUser(ctx context.Context, email, passwordHash string) (domain.User, error) {
	user := domain.User{}

	res := self.db.Where("email = ? AND password_hash = ?", email, passwordHash).WithContext(ctx).Find(&user)

	return user, res.Error
}
