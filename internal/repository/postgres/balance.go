package postgres

import "github.com/jmoiron/sqlx"

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}
