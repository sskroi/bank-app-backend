package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"gorm.io/gorm"

	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (store *PgStorage) ConvertCurrency(val decimal.Decimal,
	from, to string) (decimal.Decimal, decimal.Decimal, error) {
	var (
		res  decimal.Decimal
		rate decimal.Decimal
	)

	rows, err := store.db.Table("conversion_rates").Select("rate").Where(
		"currency_from = ? AND currency_to = ?", from, to).Rows()
	if err != nil {
		return res, rate, err
	}

	defer rows.Close()
	if !rows.Next() {
		return res, rate, fmt.Errorf("no conversion from %s to %s", from, to)
	}

	var ratePtr *decimal.Decimal
	rows.Scan(&ratePtr)
	rate = *ratePtr
	return val.Mul(rate), rate, nil
}

func (store *PgStorage) CreateTransaction(
		ctx context.Context,
		senderAcc, receiverAcc domain.Account,
		newTransaction *domain.Transaction) error {
	if newTransaction.SenderAccId == newTransaction.ReceiverAccId {
		return domain.ErrSelfAccount
	}
	if newTransaction.Sent.LessThanOrEqual(decimal.NewFromInt(0)) {
		return domain.ErrInvalidAmount
	}

	senderCurrency, receiverCurrency := senderAcc.Currency, receiverAcc.Currency
	var err error

	if senderCurrency != receiverCurrency {
		newTransaction.Received, newTransaction.ConversionRate, err = store.ConvertCurrency(
			newTransaction.Sent, senderCurrency, receiverCurrency)
		if err != nil {
			return err
		}
		newTransaction.IsConversion = true
	} else {
		newTransaction.Received = newTransaction.Sent
	}

	if err := store.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(
			&domain.Account{ID: newTransaction.SenderAccId}).WithContext(
			ctx).UpdateColumn(
			"balance", gorm.Expr("balance - ?", newTransaction.Sent))
		if res.Error != nil {
			var pgErr *pgconn.PgError
			// 23514 == check_violation error
			if errors.As(res.Error, &pgErr) && pgErr.Code == "23514" {
				return domain.ErrNegativeSenderBalance
			}
			return res.Error
		}
		if res.RowsAffected == 0 {
			return domain.ErrSenderBalance
		}

		res = tx.Model(
			&domain.Account{ID: newTransaction.ReceiverAccId}).WithContext(
			ctx).UpdateColumn(
			"balance", gorm.Expr("balance + ?", newTransaction.Received))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return domain.ErrReceiverBalance
		}

		res = store.db.Omit("id", "public_id").WithContext(
			ctx).Create(&newTransaction)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return domain.ErrUnknownTransaction
		}

		return nil
	}); err != nil {
		return err
	}

	store.db.Select("public_id").Where(
		"id = ?", newTransaction.ID).WithContext(ctx).Find(&newTransaction)

	return nil
}
