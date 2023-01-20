package handlers

import (
	_ "github.com/SerjLeo/mlf_backend/docs"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(user *models.CreateUserInput) (string, error)
	CreateUserByEmail(email string) (string, error)
	SignIn(email, password string) (string, error)
	CheckUserToken(token string) (int, error)
	SendTestEmail() error
}

type TransactionService interface {
	CreateTransaction(userId int, input *models.CreateTransactionInput) (*models.Transaction, error)
	UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error)
	DeleteTransaction(userId, transactionId int) error
	GetTransactions(userId int) ([]models.Transaction, error)
	GetTransactionById(userId, transactionId int) (models.Transaction, error)
	AttachCategory(userId int, transactionId, categoryId int) error
	DetachCategory(userId int, transactionId, categoryId int) error
}

type CategoryService interface {
	GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (models.Category, error)
	CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error)
	DeleteCategory(userId, categoryId int) error
}

type ProfileService interface {
	GetUserProfile(userId int) (*models.FullProfile, error)
	UpdateProfile(input *models.UpdateProfileInput, userId int) (*models.FullProfile, error)
}

type Service interface {
	UserService
	TransactionService
	CategoryService
	ProfileService
}

type Handler interface {
	InitRoutes() *gin.Engine
}
