package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID             uint             `gorm:"column:id;primaryKey"`
	PublicId       uuid.UUID        `gorm:"column:public_id"`
	SenderAccId    uint             `gorm:"column:sender_account_id"`
	ReceiverAccId  uint             `gorm:"column:receiver_account_id"`
	Sent		   decimal.Decimal  `gorm:"column:sent"`
	Received	   decimal.Decimal  `gorm:"column:received"`
	IsConversion   bool             `gorm:"column:is_conversion"`
	ConversionRate decimal.Decimal	`gorm:"column:conversion_rate"`
	Dt			   time.Time		`gorm:"column:dt"`
}

type TransactionExtended struct {
	PublicId			uuid.UUID		`gorm:"column:public_id"`
	SenderAccNumber		uuid.UUID		`gorm:"column:sender_account_number"`
	SenderId			uint			`gorm:"column:sender_id"`
	ReceiverAccNumber	uuid.UUID		`gorm:"column:receiver_account_number"`
	ReceiverId			uint			`gorm:"column:receiver_id"`
	Sent				decimal.Decimal	`gorm:"column:sent"`
	Received			decimal.Decimal	`gorm:"column:received"`
	IsConversion		bool			`gorm:"column:is_conversion"`
	ConversionRate		decimal.Decimal	`gorm:"column:conversion_rate"`
	Dt					time.Time		`gorm:"column:dt"`
}

var (
	ErrSelfAccount = errors.New(
		"transaction sender and receiver accounts are the same")
	ErrInvalidAmount = errors.New("invalid transaction amount")
	ErrUnknownSender = errors.New("unknown transaction sender")
	ErrUnknownReceiver = errors.New("unknown transaction receiver")
	ErrNegativeSenderBalance = errors.New("transaction sender balance overdraft")
	ErrSenderBalance = errors.New("could not update transaction sender balance")
	ErrReceiverBalance = errors.New("could not update transaction receiver balance")
	ErrUnknownTransaction = errors.New("could not create transaction")
	ErrSenderAccountClose = errors.New("sender account is close")
	ErrReceiverAccountClose = errors.New("receiver account is close")
)
