package repository

import (
	"EEP/e-wallets/model"
	"database/sql"
)

type UsersAccountRepository interface {
	Save(userAccount model.UsersAccount) error
}

type usersAccountRepository struct {
	db *sql.DB
}

func (u *usersAccountRepository) Save(userAccount model.UsersAccount) error {
	_, err := u.db.Exec(`INSERT INTO users_account(id, username, password, email, phone_number, AccountStatus, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		userAccount.Id,
		userAccount.UserName,
		userAccount.Password,
		userAccount.Email,
		userAccount.PhoneNumber,
		userAccount.AccountStatus,
		userAccount.CreatedAt,
		userAccount.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewUsersAccountRepository(db *sql.DB) UsersAccountRepository {
	return &usersAccountRepository{
		db: db,
	}
}
