package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID             uint             `gorm:"column:id;primaryKey"`
	PublicId       uuid.UUID        `gorm:"column:public_id"`
	SenderAccId    uint             `gorm:"column:sender_account_id"`
	ReceiverAccId  uint             `gorm:"column:receiver_account_id"`
	Amount         decimal.Decimal  `gorm:"column:amount"`
	IsConversion   bool             `gorm:"column:is_conversion"`
	ConversionRate *decimal.Decimal `gorm:"column:conversion_rate"`
}
