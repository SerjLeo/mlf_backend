package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetUserCategories(userId int) ([]models.Category, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1
	`, categoryTable)

	categories := []models.Category{}

	err := r.db.Select(&categories, query, userId)
	return categories, err
}

func (r *CategoryPostgres) GetUserCategoryById(userId, categoryId int) (models.Category, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1 AND category_id=$2
	`, categoryTable)

	category := models.Category{}

	err := r.db.Get(&category, query, userId, categoryId)
	return category, err
}

func (r *CategoryPostgres) CreateCategory(userId int, input models.CreateCategoryInput) (models.Category, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (name, color, user_id)
		VALUES($1, $2, $3)
		RETURNING *
	`, categoryTable)

	category := models.Category{}

	row := r.db.QueryRow(query, input.Name, input.Color, userId)
	err := row.Scan(&category.CategoryId, &category.UserId, &category.Name, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	return category, err
}

func (r *CategoryPostgres) UpdateCategory(userId, categoryId int, input models.Category) (models.Category, error) {
	return models.Category{}, nil
}