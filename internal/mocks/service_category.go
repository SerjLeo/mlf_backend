package mocks

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
)

func (o *ServiceMock) GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error) {
	return []models.Category{}, nil
}
func (o *ServiceMock) GetUserCategoryById(userId, categoryId int) (models.Category, error) {
	return models.Category{}, nil
}
func (o *ServiceMock) CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error) {
	return models.Category{}, nil
}
func (o *ServiceMock) UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error) {
	return models.Category{}, nil
}
func (o *ServiceMock) DeleteCategory(userId, categoryId int) error {
	return nil
}
