package services

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/imdario/mergo"
	"time"
)

type TransactionService struct {
	repo *Repository
}

func NewTransactionService(repo *Repository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userId int, input *models.CreateTransactionInput) (*models.Transaction, error) {
	if input.CurrencyId == 0 {
		currency, err := s.repo.CurrencyRepo.GetUsersCurrency(userId)
		fmt.Printf("%+v", currency)
		if err != nil {
			return nil, err
		}
		input.CurrencyId = currency.CurrencyId
	}
	return s.repo.TransactionRepo.CreateTransactionWithCategories(userId, *input)
}

func (s *TransactionService) UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error) {
	oldTransaction, err := s.GetTransactionById(userId, transactionId)
	if err != nil {
		return models.Transaction{}, err
	}
	mergo.Merge(&input, oldTransaction)
	input.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.TransactionRepo.UpdateTransaction(userId, transactionId, *input)
}

func (s *TransactionService) DeleteTransaction(userId, transactionId int) error {
	return s.repo.TransactionRepo.DeleteTransaction(userId, transactionId)
}

func (s *TransactionService) GetTransactions(userId int) ([]models.Transaction, error) {
	return s.repo.TransactionRepo.GetTransactions(userId)
}

func (s *TransactionService) GetTransactionById(userId, transactionId int) (models.Transaction, error) {
	transaction, err := s.repo.TransactionRepo.GetTransactionById(userId, transactionId)
	if err != nil {
		return transaction, err
	}
	categories, err := s.repo.TransactionRepo.GetTransactionCategories(userId, transactionId)
	if err != nil {
		return transaction, err
	}
	transaction.Categories = categories

	return transaction, nil
}

func (s *TransactionService) AttachCategory(userId int, transactionId, categoryId int) error {
	return s.repo.TransactionRepo.AttachCategory(userId, transactionId, categoryId)
}

func (s *TransactionService) DetachCategory(userId int, transactionId, categoryId int) error {
	return s.repo.TransactionRepo.DetachCategory(userId, transactionId, categoryId)
}
