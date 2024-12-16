package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID           uint      `gorm:"column:id;primaryKey" json:"-"`
	PublicId     uuid.UUID `gorm:"column:public_id" json:"publicId"`
	Email        string    `gorm:"column:email" json:"email"`
	PasswordHash string    `gorm:"column:password_hash" json:"-"`
	Passport     string    `gorm:"column:passport" json:"passport"`
	Name         string    `gorm:"column:name" json:"name"`
	Surname      string    `gorm:"column:surname" json:"surname"`
	Patronymic   *string   `gorm:"column:patronymic" json:"patronymic"`
	IsInactive   bool      `gorm:"column:is_inactive" json:"isInactive"`
}

var (
	ErrUserAlreadyExists       = errors.New("user with such email already exists")
	ErrUserDeleted			   = errors.New("user with this public id is no longer present")
	ErrInvalidLoginCredentials = errors.New("invalid login credentials")
)
