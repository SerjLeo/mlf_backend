package repository

import (
	postgres "github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type User interface {
}

type Transaction interface {
}

type Category interface {
}

type Repository struct {
	User
	Transaction
	Category
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        postgres.NewUserPostgres(db),
		Category:    postgres.NewCategoryPostgres(db),
		Transaction: postgres.NewTransactionPostgres(db),
	}
}
