package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/password"
)

type User interface {
	Create(user models.User) (string, error)
	CreateByEmail(email string) (string, error)
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
	Repo          *repository.Repository
	TokenManager  auth.TokenManager
	HashGenerator password.HashGenerator
}

func NewService(deps ServiceDependencies) *Service {
	return &Service{
		Category:    NewCategoryService(deps.Repo),
		User:        NewUserService(deps.Repo, deps.TokenManager, deps.HashGenerator),
		Transaction: NewTransactionService(deps.Repo),
	}
}
