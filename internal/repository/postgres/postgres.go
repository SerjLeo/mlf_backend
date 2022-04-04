package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable                = "users"
	categoryTable            = "category"
	transactionTable         = "transaction"
	transactionCategoryTable = "transaction_category"
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
