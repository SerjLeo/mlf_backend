package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable                = "users"
	profileTable             = "profile"
	accountTable             = "account"
	balanceTable             = "balance"
	accountBalanceTable      = "account_balance"
	categoryTable            = "category"
	transactionTable         = "transaction"
	transactionCategoryTable = "transaction_category"
	currencyTable            = "currency"
	defaultPerPage           = 20
)

func NewPostgresDB(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func handlePagination(query string, params models.PaginationParams) string {
	perPage := defaultPerPage
	if params.PerPage != 0 {
		perPage = params.PerPage
	}
	query = fmt.Sprintf("%s LIMIT %d", query, perPage)

	if params.Page != 0 {
		query = fmt.Sprintf("%s OFFSET %d", query, perPage*(params.Page-1))
	}

	return query
}
