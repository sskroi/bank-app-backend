package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	PublicId     uuid.UUID `gorm:"column:public_id"`
	Email        string    `gorm:"column:email"`
	PasswordHash string    `gorm:"column:password_hash"`
	Passport     string    `gorm:"column:passport"`
	Name         string    `gorm:"column:name"`
	Surname      string    `gorm:"column:surname"`
	Patronymic   *string   `gorm:"column:patronymic"`
	IsInactive   bool      `gorm:"column:is_inactive"`
}

var (
	ErrUserAlreadyExists       = errors.New("user with such email already exists")
	ErrUnknownUserPubId		   = errors.New("no user with this public id")
	ErrInvalidLoginCredentials = errors.New("invalid login credentials")
)
