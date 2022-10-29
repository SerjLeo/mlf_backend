package handlers

import (
	_ "github.com/SerjLeo/mlf_backend/docs"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User interface {
	Create(user models.User) (string, error)
	CreateUserByEmail(email string) (string, error)
	SignIn(email, password string) (string, error)
	CheckUserToken(token string) (int, error)
	GetUserProfile(userId int) (models.User, error)
	SendTestEmail() error
}

type Transaction interface {
	CreateTransaction(userId int, input *models.CreateTransactionInput) (models.Transaction, error)
	UpdateTransaction(userId, transactionId int, input *models.Transaction) (models.Transaction, error)
	DeleteTransaction(userId, transactionId int) error
	GetTransactions(userId int) ([]models.Transaction, error)
	GetTransactionById(userId, transactionId int) (models.Transaction, error)
	AttachCategory(userId int, transactionId, categoryId int) error
	DetachCategory(userId int, transactionId, categoryId int) error
}

type Category interface {
	GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error)
	GetUserCategoryById(userId, categoryId int) (models.Category, error)
	CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error)
	DeleteCategory(userId, categoryId int) error
}

type Service interface {
	User
	Transaction
	Category
}

type Handler interface {
	InitRoutes() *gin.Engine
}

type BotHandler interface {
	HandleMessage(msg *tgbotapi.Message) error
}
