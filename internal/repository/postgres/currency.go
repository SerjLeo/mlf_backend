package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CurrencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgres(db *sqlx.DB) *CurrencyPostgres {
	return &CurrencyPostgres{db: db}
}

func (r *CurrencyPostgres) GetCurrencyById(currencyId int) (models.Currency, error) {
	var currency models.Currency
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE currency_id=$1
	`, currencyTable)

	err := r.db.Get(&currency, query, currencyId)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return currency, models.CurrencyNotFound
	}

	return currency, err
}

func (r *CurrencyPostgres) GetCurrencyList() ([]models.Currency, error) {

	query := fmt.Sprintf(`
		SELECT * FROM %s
	`, currencyTable)

	currencies := []models.Currency{}

	err := r.db.Select(&currencies, query)
	return currencies, err
}

func (r *CurrencyPostgres) GetUsersCurrency(userId int) (*models.Currency, error) {

	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE currency_id=(
			SELECT currency_id FROM %s WHERE user_id=$1
		)
	`, currencyTable, userTable)

	currency := &models.Currency{}

	err := r.db.Get(currency, query, userId)
	return currency, err
}
