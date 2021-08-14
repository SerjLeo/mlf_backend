package services

import "github.com/SerjLeo/mlf_backend/internal/repository"

type User interface {
}

type Transaction interface {
}

type Category interface {
}

type Service struct {
	User
	Transaction
	Category
}

type ServiceDependencies struct {
	repo *repository.Repository
}

func NewService(deps ServiceDependencies) *Service {
	return &Service{
		Category:    NewCategoryService(deps.repo),
		User:        NewUserService(deps.repo),
		Transaction: NewTransactionService(deps.repo),
	}
}
