package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID       uint            `gorm:"column:id;primaryKey" json:"-"`
	Number   uuid.UUID       `gorm:"column:number"   json:"number"`
	OwnerId  uint            `gorm:"column:owner_id" json:"-"`
	Balance  decimal.Decimal `gorm:"column:balance"  json:"balance"`
	Currency string          `gorm:"column:currency" json:"currency"`
	IsClose  bool            `gorm:"column:is_close" json:"isClose"`
}

var (
	ErrUnknownCurrency = errors.New("unknown currency")
	ErrUnknownAccount = errors.New("unknown account")
	ErrAlreadyClose = errors.New("account is already close")
	ErrClose = errors.New("could not close account")
)
