package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) CreateAccount(name string, currencyId, userId int) (int, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (name, user_id)
		VALUES ($1, $2)
		RETURNING account_id
	`, accountTable)

	row := tx.QueryRow(query, name, userId)
	var accountId int
	if err = row.Scan(&accountId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(`
		INSERT INTO %s (currency_id, user_id)
		VALUES ($1, $2)
		RETURNING balance_id
	`, balanceTable)

	row = tx.QueryRow(query, currencyId, userId)
	var balanceId int
	if err = row.Scan(&balanceId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(`
		INSERT INTO %s (account_id, balance_id)
		VALUES ($1, $2)
	`, accountBalanceTable)

	_, err = tx.Exec(query, accountId, balanceId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return accountId, nil
}

func (r *AccountPostgres) GetAccountById(accountId, userId int) (*models.Account, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1 AND account_id=$2
	`, accountTable)

	account := &models.Account{}
	err := r.db.Get(&account, query, userId, accountId)

	return account, err
}

func (r *AccountPostgres) GetUsersAccounts(userId int, pagination models.PaginationParams) ([]models.Account, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1
	`, accountTable)

	accounts := []models.Account{}

	err := r.db.Select(&accounts, handlePagination(query, pagination), userId)

	return accounts, err
}

func (r *AccountPostgres) UpdateAccount(accountId, userId int, input *models.Account) error {
	return nil
}

func (r *AccountPostgres) DeleteAccount(accountId, userId int) error {
	return nil
}
