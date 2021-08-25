package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User) (int, error) {
	var id int
	query := `
		INSERT INTO users (name, email, hashed_pass)
		VALUES ($1, $2, $3)
		RETURNING user_id`

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return id, err
	}

	return id, nil

}
