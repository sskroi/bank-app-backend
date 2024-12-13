package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
)

func (store *PgStorage) CreateAccount(
		ctx context.Context, newAccount *domain.Account) error {
	res := store.db.Omit("id", "number").WithContext(ctx).Create(newAccount)
	if res.Error != nil {
		return res.Error
	}

	// this can be omitted
	store.db.Select("number").Where(
		"id = ?", newAccount.ID).WithContext(ctx).Find(newAccount)

	return nil
}

func (store *PgStorage) GetUserAccounts(ctx context.Context, userId uint,
		offset, limit int) ([]domain.Account, error) {
	var accounts []domain.Account

	if res := store.db.Where("owner_id = ?", userId).WithContext(
			ctx).Offset(offset).Limit(limit).Find(&accounts); res.Error != nil {
		return accounts, res.Error
	}

	return accounts, nil
}

func (store *PgStorage) GetAccountCurrency(
		ctx context.Context, accId uint,
		notFoundErr error) (string, error) {
	var acc domain.Account

	res := store.db.Select("currency").Where(
		"id = ?", accId).WithContext(ctx).Find(&acc)
	err := res.Error
	if err == nil && res.RowsAffected == 0 {
		if notFoundErr != nil {
			err = notFoundErr
		} else {
			err = domain.ErrUnknownAccount
		}
	}
	return acc.Currency, err
}
