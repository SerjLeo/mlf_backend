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
	Repo *repository.Repository
}

func NewService(deps ServiceDependencies) *Service {
	return &Service{
		Category:    NewCategoryService(deps.Repo),
		User:        NewUserService(deps.Repo),
		Transaction: NewTransactionService(deps.Repo),
	}
}
