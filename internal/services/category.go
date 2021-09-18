package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
)

type CategoryService struct {
	repo *repository.Repository
}

func NewCategoryService(repo *repository.Repository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetUserCategories(userId int) ([]models.Category, error) {
	return s.repo.Category.GetUserCategories(userId)
}

func (s *CategoryService) GetUserCategoryById(userId, categoryId int) (models.Category, error) {
	return s.repo.Category.GetUserCategoryById(userId, categoryId)
}

func (s *CategoryService) CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error) {
	return s.repo.Category.CreateCategory(userId, input)
}
