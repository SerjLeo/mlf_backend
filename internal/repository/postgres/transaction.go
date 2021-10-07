package postgres

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) CreateTransaction(userId int, input models.Transaction, categoriesIds []int) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (r *TransactionPostgres) UpdateTransaction(userId, transactionId int, input models.Transaction) (models.Transaction, error){
	return models.Transaction{}, nil
}

func (r *TransactionPostgres) DeleteTransaction(userId, transactionId int) (int, error) {
	return 0, nil
}

func (r *TransactionPostgres) GetTransactions(userId int) ([]models.Transaction, error){
	return []models.Transaction{}, nil
}

func (r *TransactionPostgres) GetTransactionById(userId, transactionId int) (models.Transaction, error){
	return models.Transaction{}, nil
}

func (r *TransactionPostgres) AttachCategory(userId, transactionId int, categoriesIds []int) (models.Transaction, error){
	return models.Transaction{}, nil
}

func (r *TransactionPostgres) DetachCategory(userId, transactionId int, categoriesIds []int) (models.Transaction, error){
	return models.Transaction{}, nil
}