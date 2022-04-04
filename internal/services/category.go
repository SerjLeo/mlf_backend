package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"time"
)

type CategoryService struct {
	repo   *repository.Repository
	colors colors.ColorManager
}

func NewCategoryService(repo *repository.Repository, colors colors.ColorManager) *CategoryService {
	return &CategoryService{repo: repo, colors: colors}
}

func (s *CategoryService) GetUserCategories(userId int) ([]models.Category, error) {
	return s.repo.Category.GetUserCategories(userId)
}

func (s *CategoryService) GetUserCategoryById(userId, categoryId int) (models.Category, error) {
	return s.repo.Category.GetUserCategoryById(userId, categoryId)
}

func (s *CategoryService) CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error) {
	if !s.colors.IsHEX(input.Color) {
		input.Color = s.colors.GenerateHex()
	}
	return s.repo.Category.CreateCategory(userId, input)
}

func (s *CategoryService) UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error) {
	oldCategory, err := s.GetUserCategoryById(userId, categoryId)
	if err != nil {
		return models.Category{}, err
	}
	mergo.Merge(&input, oldCategory)
	input.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.Category.UpdateCategory(userId, categoryId, input)
}

func (s *CategoryService) DeleteCategory(userId, categoryId int) error {
	_, err := s.GetUserCategoryById(userId, categoryId)
	if err != nil {
		return errors.Wrap(err, "error while checking category existence")
	}
	return s.repo.DeleteCategory(userId, categoryId)
}
