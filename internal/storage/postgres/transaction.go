package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
	"fmt"

	"github.com/google/uuid"
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
	if err := rows.Scan(&ratePtr); err != nil {
		return res, rate, err
	}
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

		res = store.db.Omit("id", "public_id", "dt").WithContext(
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

	store.db.Select("public_id", "dt").Where(
		"id = ?", newTransaction.ID).WithContext(ctx).Find(&newTransaction)

	fmt.Printf("New transaction: %+v", newTransaction)

	return nil
}

// SELECT public_id, sender_acc.number AS sender_account_number, sender_acc.owner_id AS sender_id, receiver_acc.number AS receiver_account_number, receiver_acc.owner_id AS receiver_id, sent, received, is_conversion, conversion_rate, dt FROM transactions JOIN accounts AS sender_acc ON sender_acc.id = transactions.sender_account_id JOIN accounts AS receiver_acc ON receiver_acc.id = transactions.receiver_account_id;

func (store *PgStorage) GetUserTransactions(
		ctx context.Context, userId uint,
		accountNumber *uuid.UUID, offset, limit int) ([]domain.TransactionExtended, error) {
	var transactions []domain.TransactionExtended

	res := store.db.Model(&domain.Transaction{}).Select(
		"public_id",
		fmt.Sprintf("CASE WHEN sender_acc.owner_id = %d THEN sender_acc.number ELSE NULL END AS sender_account_number", userId),
		"sender_acc.owner_id AS sender_id",
		"receiver_acc.number AS receiver_account_number",
		"receiver_acc.owner_id AS receiver_id", "sent",
		"sender_acc.currency AS sent_currency", "received",
		"receiver_acc.currency AS received_currency", "is_conversion", "conversion_rate",
		fmt.Sprintf("CASE WHEN sender_acc.owner_id = %d AND receiver_acc.owner_id = %d THEN 0 WHEN sender_acc.owner_id = %d THEN -1 ELSE 1 END AS direction",
					userId, userId, userId),
		"sender_acc.owner_id = receiver_acc.owner_id AS same_owner", "dt",
		).Joins("JOIN accounts AS sender_acc ON sender_acc.id = transactions.sender_account_id",
		).Joins("JOIN accounts AS receiver_acc ON receiver_acc.id = transactions.receiver_account_id",
		).Where("sender_acc.owner_id = ? OR receiver_acc.owner_id = ?",
				userId, userId).Order("transactions.id DESC")
	if accountNumber != nil {
		res = res.Where(
			"receiver_acc.number = ? OR sender_acc.number = ? AND sender_acc.owner_id = ?",
			accountNumber, accountNumber, userId)
	}
	res = res.Offset(offset).Limit(limit).Find(&transactions)
	
	return transactions, res.Error
}
