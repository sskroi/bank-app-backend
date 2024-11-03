package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID       uint            `gorm:"column:id;primaryKey" json:"-"`
	Number   uuid.UUID       `gorm:"column:number" json:"number"`
	OwnerId  uint            `gorm:"column:owner_id" json:"-"`
	Balance  decimal.Decimal `gorm:"column:balance" json:"balance"`
	Currency string          `gorm:"column:currency" json:"currency"`
	IsClose  bool            `gorm:"column:is_close" json:"is_close"`
}
