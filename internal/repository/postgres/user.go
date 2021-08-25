package postgres

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
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
		if strings.Contains(err.Error(), "duplicate key value") {
			return id, models.UserAlreadyExist
		}
		return id, err
	}

	return id, nil
}

func (r *UserPostgres) GetUser(email, passHash string) (models.User, error) {
	var user models.User
	query := `
		SELECT * FROM users
		WHERE email=$1 AND hashed_pass=$2`

	err := r.db.Get(&user, query, email, passHash)
	if strings.Contains(err.Error(), "no rows") {
		return user, models.EmailOrPassNotMatch
	}

	return user, err
}
