package repository

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	postgres "github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) (int, error)
	GetUser(email, passHash string) (models.User, error)
}

type Transaction interface {
}

type Category interface {
	GetUserCategories(userId int) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (models.Category, error)
	CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error)
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
