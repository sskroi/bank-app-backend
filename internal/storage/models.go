package storage

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	ID           uint
	PublicId     uuid.UUID
	PasswordHash string
	Email        string
	Name         string
	Surname      string
	Patronymic   *string
	Passport     string
	IsInactive   bool
}

type Account struct {
	ID            uint
	AccountNumber uuid.UUID
	OwnerId       uint
	Balance       decimal.Decimal
	Currency      string
	IsClose       bool
}

type Transaction struct {
	ID             uint
	PublicId       uuid.UUID
	SenderAccId    uint
	ReceiverAccId  uint
	Amount         decimal.Decimal
	IsConversion   bool
	ConversionRate *decimal.Decimal
}
