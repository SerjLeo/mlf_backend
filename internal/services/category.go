package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/pkg/errors"
	"time"
)

type CategoryService struct {
	repo   *Repository
	colors colors.ColorManager
}

func NewCategoryService(repo *Repository, colors colors.ColorManager) *CategoryService {
	return &CategoryService{repo: repo, colors: colors}
}

func (s *CategoryService) GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error) {
	return s.repo.CategoryRepo.GetUserCategories(userId, pagination)
}

func (s *CategoryService) GetUserCategoryById(userId, categoryId int) (*models.Category, error) {
	return s.repo.CategoryRepo.GetUserCategoryById(userId, categoryId)
}

func (s *CategoryService) CreateCategory(userId int, input *models.CreateCategoryInput) (*models.Category, error) {
	if !s.colors.IsHEX(input.Color) {
		input.Color = s.colors.GenerateHex()
	}
	return s.repo.CategoryRepo.CreateCategory(userId, input)
}

func (s *CategoryService) UpdateCategory(userId, categoryId int, input *models.UpdateCategoryInput) (*models.Category, error) {
	input.UpdatedAt = time.Now().Format(time.RFC3339)
	err := s.repo.CategoryRepo.UpdateCategory(userId, categoryId, input)
	if err != nil {
		return nil, err
	}
	return s.GetUserCategoryById(userId, categoryId)
}

func (s *CategoryService) DeleteCategory(userId, categoryId int) error {
	_, err := s.GetUserCategoryById(userId, categoryId)
	if err != nil {
		return errors.Wrap(err, "error while checking category existence")
	}
	return s.repo.DeleteCategory(userId, categoryId)
}
