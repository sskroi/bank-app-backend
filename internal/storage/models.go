package storage

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

type Account struct {
	ID            uint            `gorm:"column:id;primaryKey"`
	AccountNumber uuid.UUID       `gorm:"column:account_number"`
	OwnerId       uint            `gorm:"column:owner_id"`
	Balance       decimal.Decimal `gorm:"column:balance"`
	Currency      string          `gorm:"column:currency"`
	IsClose       bool            `gorm:"column:is_close"`
}

type Transaction struct {
	ID             uint             `gorm:"column:id;primaryKey"`
	PublicId       uuid.UUID        `gorm:"column:public_id"`
	SenderAccId    uint             `gorm:"column:sender_account_id"`
	ReceiverAccId  uint             `gorm:"column:receiver_account_id"`
	Amount         decimal.Decimal  `gorm:"column:amount"`
	IsConversion   bool             `gorm:"column:is_conversion"`
	ConversionRate *decimal.Decimal `gorm:"column:conversion_rate"`
}
