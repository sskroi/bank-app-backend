package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID       uint            `gorm:"column:id;primaryKey"`
	Number   uuid.UUID       `gorm:"column:number"`
	OwnerId  uint            `gorm:"column:owner_id"`
	Balance  decimal.Decimal `gorm:"column:balance"`
	Currency string          `gorm:"column:currency"`
	IsClose  bool            `gorm:"column:is_close"`
}
