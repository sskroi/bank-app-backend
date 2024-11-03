package postgres

import (
	"bank-app-backend/internal/domain"
	"context"
	// "errors"
	// "github.com/jackc/pgx/v5/pgconn"
)

func (store *PgStorage) GetUserAccounts(ctx context.Context, uid uint,
		offset, limit int) ([]domain.Account, error) {
	if limit == 0 || limit > 100 {
		limit = 100
	}

	var accounts []domain.Account

	if res := store.db.Where("owner_id = ?", uid).WithContext(
			ctx).Offset(offset).Limit(limit).Find(&accounts); res.Error != nil {
		return accounts, res.Error
	}

	return accounts, nil
}
