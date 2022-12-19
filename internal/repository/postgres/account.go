package postgres

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) CreateAccount(name string, userId int) (int, error) {
	return 0, nil
}

func (r *AccountPostgres) GetAccountById(accountId, userId int) (*models.Account, error) {
	return nil, nil
}

func (r *AccountPostgres) GetUsersAccounts(userId int, pagination models.PaginationParams) (*models.Account, error) {
	return nil, nil
}

func (r *AccountPostgres) UpdateAccount(accountId, userId int, input *models.Account) error {
	return nil
}

func (r *AccountPostgres) DeleteAccount(accountId, userId int) error {
	return nil
}
