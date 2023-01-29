package repository

import (
	postgres "github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/jmoiron/sqlx"
)

func NewPostgresRepository(db *sqlx.DB) *services.Repository {
	return &services.Repository{
		UserRepo:        postgres.NewUserPostgres(db),
		CategoryRepo:    postgres.NewCategoryPostgres(db),
		TransactionRepo: postgres.NewTransactionPostgres(db),
		CurrencyRepo:    postgres.NewCurrencyPostgres(db),
		AccountRepo:     postgres.NewAccountPostgres(db),
		ProfileRepo:     postgres.NewProfilePostgres(db),
		BalanceRepo:     postgres.NewBalancePostgres(db),
	}
}
