package services

import (
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/SerjLeo/mlf_backend/pkg/email"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
)

type ServiceDependencies struct {
	Repo            *repository.Repository
	TokenManager    auth.TokenManager
	HashGenerator   password.HashGenerator
	MailManager     email.MailManager
	TemplateManager templates.TemplateManager
	Cache           *cache.Cache
	ColorManager    colors.ColorManager
}

type AppService struct {
	CategoryService
	UserService
	TransactionService
}

func NewService(deps ServiceDependencies) *AppService {
	return &AppService{
		*NewCategoryService(deps.Repo, deps.ColorManager),
		*NewUserService(deps.Repo, deps.TokenManager, deps.HashGenerator, deps.MailManager, deps.TemplateManager, deps.Cache),
		*NewTransactionService(deps.Repo),
	}
}
