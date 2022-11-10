package mocks

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
)

func (o *ServiceMock) CreateTransaction(userId int, input *models.CreateTransactionInput) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}
func (o *ServiceMock) UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error) {
	return models.Transaction{}, nil
}
func (o *ServiceMock) DeleteTransaction(userId, transactionId int) error {
	return nil
}
func (o *ServiceMock) GetTransactions(userId int) ([]models.Transaction, error) {
	return []models.Transaction{}, nil
}
func (o *ServiceMock) GetTransactionById(userId, transactionId int) (models.Transaction, error) {
	return models.Transaction{}, nil
}
func (o *ServiceMock) AttachCategory(userId int, transactionId, categoryId int) error {
	return nil
}
func (o *ServiceMock) DetachCategory(userId int, transactionId, categoryId int) error {
	return nil
}
