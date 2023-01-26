package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
)

type UserRepo interface {
	CreateUser(input *models.CreateUserInput) (int, error)
	AuthenticateUser(email, passHash string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)
	ChangePassword(userId int, password string) error
}

type TransactionRepo interface {
	CreateTransaction(userId int, input models.CreateTransactionInput) (*models.Transaction, error)
	CreateTransactionWithCategories(userId int, input models.CreateTransactionInput) (*models.Transaction, error)
	UpdateTransaction(userId, transactionId int, input models.Transaction) (models.Transaction, error)
	DeleteTransaction(userId, transactionId int) error
	GetTransactions(userId int) ([]models.Transaction, error)
	GetTransactionById(userId, transactionId int) (models.Transaction, error)
	AttachCategory(userId, transactionId int, categoryId int) error
	DetachCategory(userId, transactionId int, categoryId int) error
	GetTransactionCategories(userId, transactionId int) ([]models.Category, error)
}

type CategoryRepo interface {
	GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (*models.Category, error)
	CreateCategory(userId int, input *models.CreateCategoryInput) (*models.Category, error)
	UpdateCategory(userId, categoryId int, input *models.UpdateCategoryInput) (*models.Category, error)
	DeleteCategory(userId, categoryId int) error
}

type CurrencyRepo interface {
	GetCurrencyList() ([]models.Currency, error)
	GetCurrencyById(currencyId int) (models.Currency, error)
	GetUsersCurrency(userId int) (*models.Currency, error)
}

type AccountRepo interface {
	CreateAccount(name string, currencyId, userId int) (int, error)
	GetAccountById(accountId, userId int) (*models.Account, error)
	GetUsersAccounts(userId int, pagination models.PaginationParams) ([]models.Account, error)
	UpdateAccount(accountId, userId int, input *models.UpdateAccountInput) error
	SoftDeleteAccount(accountId, userId int) error
}

type ProfileRepo interface {
	CreateProfile(userId int, name string) (int, error)
	UpdateProfile(input *models.UpdateProfileInput, userId int) error
	GetUserProfile(userId int) (*models.FullProfile, error)
	DeleteProfile(userId, profile int) error
}

type BalanceRepo interface {
	GetAccountBalances(userId, accountId int) (*[]models.Balance, error)
	CreateBalance(userId, accountId, currencyId int) (int, error)
	GetUserCurrencyBalanceAmount(userId, currencyId int) (*models.BalanceOfCurrency, error)
	GetUserBalancesAmount(userId int) ([]models.BalanceOfCurrency, error)
}

type Repository struct {
	UserRepo
	TransactionRepo
	CategoryRepo
	CurrencyRepo
	AccountRepo
	ProfileRepo
	BalanceRepo
}

type MailManager interface {
	SendEmail(to string, subject string, body string) error
}

type ServiceDependencies struct {
	Repo            *Repository
	TokenManager    auth.TokenManager
	HashGenerator   password.HashGenerator
	MailManager     MailManager
	TemplateManager templates.TemplateManager
	Cache           *cache.Cache
	ColorManager    colors.ColorManager
	Env             string
}

type AppService struct {
	CategoryService
	UserService
	TransactionService
	ProfileService
	AccountService
	BalanceService
}

func NewService(deps ServiceDependencies) *AppService {
	return &AppService{
		*NewCategoryService(deps.Repo, deps.ColorManager),
		*NewUserService(deps.Repo, deps.TokenManager, deps.HashGenerator, deps.MailManager, deps.TemplateManager, deps.Cache, deps.Env),
		*NewTransactionService(deps.Repo),
		*NewProfileService(deps.Repo),
		*NewAccountService(deps.Repo),
		*NewBalanceService(deps.Repo),
	}
}
