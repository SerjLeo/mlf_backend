package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
)

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userId int, input *models.CreateTransactionInput) (models.Transaction, error) {
	return s.repo.Transaction.CreateTransactionWithCategories(userId, *input)
}

func (s *TransactionService) UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error) {
	return s.repo.Transaction.UpdateTransaction(userId, transactionId, *input)
}

func (s *TransactionService) DeleteTransaction(userId, transactionId int) (int, error) {
	return s.repo.Transaction.DeleteTransaction(userId, transactionId)
}

func (s *TransactionService) GetTransactions(userId int) ([]models.Transaction, error) {
	return s.repo.Transaction.GetTransactions(userId)
}

func (s *TransactionService) GetTransactionById(userId, transactionId int) (models.Transaction, error) {
	return s.repo.Transaction.GetTransactionById(userId, transactionId)
}

func (s *TransactionService) AttachCategory(userId int, transaction *models.Transaction, categoryId int) error {
	return s.repo.Transaction.AttachCategory(userId, transaction.TransactionId, categoryId)
}

func (s *TransactionService) DetachCategory(userId int, transaction *models.Transaction, categoryId int) error {
	return s.repo.Transaction.DetachCategory(userId, transaction.TransactionId, categoryId)
}
