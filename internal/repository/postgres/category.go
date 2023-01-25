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

func (r *CategoryPostgres) GetUserCategories(userId int, pagination models.PaginationParams) ([]models.Category, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1
	`, categoryTable)
	queryWithPagination := handlePagination(query, pagination)

	categories := []models.Category{}

	err := r.db.Select(&categories, queryWithPagination, userId)
	return categories, err
}

func (r *CategoryPostgres) GetUserCategoryById(userId, categoryId int) (*models.Category, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1 AND category_id=$2
	`, categoryTable)

	category := models.Category{}

	err := r.db.Get(&category, query, userId, categoryId)
	return &category, err
}

func (r *CategoryPostgres) CreateCategory(userId int, input *models.CreateCategoryInput) (*models.Category, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (name, color, user_id)
		VALUES($1, $2, $3)
		RETURNING *
	`, categoryTable)

	category := models.Category{}

	row := r.db.QueryRow(query, input.Name, input.Color, userId)
	err := row.Scan(&category.Id, &category.UserId, &category.Name, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	return &category, err
}

func (r *CategoryPostgres) UpdateCategory(userId, categoryId int, input *models.UpdateCategoryInput) (*models.Category, error) {
	query := fmt.Sprintf(`
		UPDATE %s
		SET name = $1, color = $2, updated_at = $3
		WHERE user_id = $4 AND category_id = $5
		RETURNING *
	`, categoryTable)

	category := models.Category{}

	row := r.db.QueryRow(query, input.Name, input.Color, input.UpdatedAt, userId, categoryId)
	err := row.Scan(&category.Id, &category.UserId, &category.Name, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	return &category, err
}

func (r *CategoryPostgres) DeleteCategory(userId, categoryId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = $1 AND category_id = $2
	`, categoryTable)

	_, err := r.db.Exec(query, userId, categoryId)
	return err
}
