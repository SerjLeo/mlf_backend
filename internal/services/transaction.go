package services

import "github.com/SerjLeo/mlf_backend/internal/repository"

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo: repo}
}
