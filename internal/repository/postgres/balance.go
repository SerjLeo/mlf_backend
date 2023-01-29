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

func (r *BalancePostgres) GetBalanceById(userId, balanceId int) (*models.Balance, error) {
	query := fmt.Sprintf(`
		SELECT bal.balance_id, amount, cur.currency_id, name as currency, bal.created_at as created_at
		FROM %s AS bal
         INNER JOIN %s AS cur ON bal.currency_id=cur.currency_id
         INNER JOIN %s AS acc_bal ON acc_bal.balance_id=bal.balance_id
		WHERE user_id=$1 AND balance_id=$2
	`, balanceTable, currencyTable, accountBalanceTable)

	balance := models.Balance{}

	err := r.db.Get(&balance, query, userId, balanceId)

	return &balance, err
}

func (r *BalancePostgres) GetBalanceByCurrencyAndAccount(userId, currencyId, accountId int) (*models.Balance, error) {
	query := fmt.Sprintf(`
		SELECT bal.balance_id, amount, cur.currency_id, name as currency, bal.created_at as created_at
		FROM %s AS bal
         INNER JOIN %s AS cur ON bal.currency_id=cur.currency_id
         INNER JOIN %s AS acc_bal ON acc_bal.balance_id=bal.balance_id
		WHERE user_id=$1 AND account_id=$2 AND currency_id=$3
	`, balanceTable, currencyTable, accountBalanceTable)

	balance := models.Balance{}

	err := r.db.Get(&balance, query, userId, accountId, currencyId)

	return &balance, err
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

func (r *BalancePostgres) GetUserBalancesAmount(userId int) ([]models.BalanceOfCurrency, error) {
	query := fmt.Sprintf(`
		SELECT cur.currency_id, cur.name as currency, SUM(bal.amount) as total, bal.user_id
		FROM %s as bal
		INNER JOIN %s as cur ON bal.currency_id=cur.currency_id
	 	WHERE bal.user_id=$1
		GROUP BY cur.currency_id, user_id
	`, balanceTable, currencyTable)

	balances := []models.BalanceOfCurrency{}

	err := r.db.Select(&balances, query, userId)

	return balances, err
}

func (r *BalancePostgres) GetUserCurrencyBalanceAmount(userId, currencyId int) (*models.BalanceOfCurrency, error) {
	query := fmt.Sprintf(`
		SELECT cur.currency_id, cur.name as currency, SUM(bal.amount) as total, bal.user_id
		FROM %s as bal
		INNER JOIN %s as cur ON bal.currency_id=cur.currency_id
		WHERE cur.currency_id=$1 AND user_id=$2
		GROUP BY cur.currency_id, user_id
	`, balanceTable, currencyTable)

	balance := &models.BalanceOfCurrency{}

	err := r.db.Get(balance, query, currencyId, userId)
	return balance, err
}

func (r *BalancePostgres) UpdateBalanceValue(userId, balanceId int, value float64) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET amount=$1
		WHERE user_id=$2 AND balance_id=$3
	`, balanceTable)

	_, err := r.db.Exec(query, value, userId, balanceId)
	return err
}

func (r *BalancePostgres) DeleteBalance(userId, balanceId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND balance_id=$2", balanceTable)

	_, err := r.db.Exec(query, userId, balanceId)
	return err
}
