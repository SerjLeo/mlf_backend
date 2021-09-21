package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/SerjLeo/mlf_backend/pkg/email"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
)

type User interface {
	Create(user models.User) (string, error)
	CreateByEmail(email string) (string, error)
	SignIn(email, password string) (string, error)
	CheckUserToken(token string) (int, error)
	SendTestEmail() error
}

type Transaction interface {
}

type Category interface {
	GetUserCategories(userId int) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (models.Category, error)
	CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error)
}

type Service struct {
	User
	Transaction
	Category
}

type ServiceDependencies struct {
	Repo            *repository.Repository
	TokenManager    auth.TokenManager
	HashGenerator   password.HashGenerator
	MailManager     email.MailManager
	TemplateManager templates.TemplateManager
	Cache           *cache.Cache
	ColorManager    colors.ColorManager
}

func NewService(deps ServiceDependencies) *Service {
	return &Service{
		Category:    NewCategoryService(deps.Repo, deps.ColorManager),
		User:        NewUserService(deps.Repo, deps.TokenManager, deps.HashGenerator, deps.MailManager, deps.TemplateManager, deps.Cache),
		Transaction: NewTransactionService(deps.Repo),
	}
}
