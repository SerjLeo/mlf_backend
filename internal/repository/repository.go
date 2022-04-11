package repository

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	postgres "github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) (int, error)
	GetUser(email, passHash string) (models.User, error)
	GetUserById(userId int) (models.User, error)
}

type Transaction interface {
	CreateTransaction(userId int, input models.CreateTransactionInput) (models.Transaction, error)
	CreateTransactionWithCategories(userId int, input models.CreateTransactionInput) (models.Transaction, error)
	UpdateTransaction(userId, transactionId int, input models.Transaction) (models.Transaction, error)
	DeleteTransaction(userId, transactionId int) (int, error)
	GetTransactions(userId int) ([]models.Transaction, error)
	GetTransactionById(userId, transactionId int) (models.Transaction, error)
	AttachCategory(userId, transactionId int, categoryId int) error
	DetachCategory(userId, transactionId int, categoryId int) error
	GetTransactionCategories(userId, transactionId int) ([]models.Category, error)
}

type Category interface {
	GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (models.Category, error)
	CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error)
	DeleteCategory(userId, categoryId int) error
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
