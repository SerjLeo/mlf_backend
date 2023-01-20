package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) CreateBalance(userId, accountId, currencyId int) (int, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (currency_id, user_id)
		VALUES ($1, $2)
		RETURNING balance_id
	`, balanceTable)

	var balanceId int
	row := tx.QueryRow(query, currencyId, userId)
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

	return balanceId, nil
}

func (r *BalancePostgres) GetAccountBalances(userId, accountId int) (*[]models.Balance, error) {
	query := fmt.Sprintf(`
		SELECT bal.balance_id, amount, cur.currency_id, name as currency, bal.created_at as created_at
		FROM %s AS bal
         INNER JOIN %s AS cur ON bal.currency_id=cur.currency_id
         INNER JOIN %s AS acc_bal ON acc_bal.balance_id=bal.balance_id
		WHERE user_id=$1 AND account_id=$2
	`, balanceTable, currencyTable, accountBalanceTable)

	balances := []models.Balance{}

	err := r.db.Select(&balances, query, userId, accountId)

	return &balances, err
}

func (r *BalancePostgres) GetUserBalanceAmount(userId, currencyId int) (int, error) {
	//query := fmt.Sprintf(`
	//	SELECT FROM %s (currency_id, user_id)
	//	VALUES ($1, $2)
	//`)
	return 0, nil
}
