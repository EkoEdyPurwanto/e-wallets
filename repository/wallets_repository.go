package repository

import (
	"EEP/e-wallets/model"
	"database/sql"
)

type WalletsRepository interface {
	Save(wallets model.Wallets) error
}

type walletsRepository struct {
	db *sql.DB
}

func (w *walletsRepository) Save(wallets model.Wallets) error {
	_, err := w.db.Exec(`INSERT INTO wallets (id, user_id, rekening_user, balance, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`,
		wallets.Id,
		wallets.UserId,
		wallets.RekeningUser,
		wallets.Balance,
		wallets.CreatedAt,
		wallets.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewWalletsRepository(db *sql.DB) WalletsRepository {
	return &walletsRepository{
		db: db,
	}
}
