package postgres

import (
	"bank-app-backend/internal/domain"
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (store *PgStorage) GetAccountByNumber(
	ctx context.Context, number uuid.UUID,
	ownerId uint, notFoundErr error) (domain.Account, error) {
	var acc domain.Account

	var tx *gorm.DB
	if ownerId != 0 {
		tx = store.db.Where("number = ? AND owner_id = ?", number, ownerId)
	} else {
		tx = store.db.Where("number = ?", number)
	}
	res := tx.WithContext(ctx).Find(&acc)
	err := res.Error
	if err == nil && res.RowsAffected == 0 {
		if notFoundErr != nil {
			err = notFoundErr
		} else {
			err = domain.ErrUnknownAccount
		}
	}
	return acc, err
}

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

func (store *PgStorage) CloseAccount(
		ctx context.Context, number uuid.UUID, ownerId uint) error {
	account, err := store.GetAccountByNumber(ctx, number, ownerId, nil)
	if err != nil {
		return err
	}

	if account.IsClose {
		return domain.ErrAlreadyClose
	}

	account.IsClose = true
	res := store.db.WithContext(ctx).Save(&account)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return domain.ErrClose
	}

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
