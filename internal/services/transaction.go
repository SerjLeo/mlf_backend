package services

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/imdario/mergo"
	"time"
)

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userId int, input *models.CreateTransactionInput) (*models.Transaction, error) {
	if input.CurrencyId == 0 {
		currency, err := s.repo.Currency.GetUsersCurrency(userId)
		fmt.Printf("%+v", currency)
		if err != nil {
			return nil, err
		}
		input.CurrencyId = currency.CurrencyId
	}
	return s.repo.Transaction.CreateTransactionWithCategories(userId, *input)
}

func (s *TransactionService) UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error) {
	oldTransaction, err := s.GetTransactionById(userId, transactionId)
	if err != nil {
		return models.Transaction{}, err
	}
	mergo.Merge(&input, oldTransaction)
	input.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.Transaction.UpdateTransaction(userId, transactionId, *input)
}

func (s *TransactionService) DeleteTransaction(userId, transactionId int) error {
	return s.repo.Transaction.DeleteTransaction(userId, transactionId)
}

func (s *TransactionService) GetTransactions(userId int) ([]models.Transaction, error) {
	return s.repo.Transaction.GetTransactions(userId)
}

func (s *TransactionService) GetTransactionById(userId, transactionId int) (models.Transaction, error) {
	transaction, err := s.repo.Transaction.GetTransactionById(userId, transactionId)
	if err != nil {
		return transaction, err
	}
	categories, err := s.repo.Transaction.GetTransactionCategories(userId, transactionId)
	if err != nil {
		return transaction, err
	}
	transaction.Categories = categories

	return transaction, nil
}

func (s *TransactionService) AttachCategory(userId int, transactionId, categoryId int) error {
	return s.repo.Transaction.AttachCategory(userId, transactionId, categoryId)
}

func (s *TransactionService) DetachCategory(userId int, transactionId, categoryId int) error {
	return s.repo.Transaction.DetachCategory(userId, transactionId, categoryId)
}
