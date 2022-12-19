package postgres

import (
	"fmt"
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

func (r *UserPostgres) CreateUser(input *models.CreateUserInput) (int, error) {
	var id int

	tx, err := r.db.Beginx()
	if err != nil {
		return id, err
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (email, hashed_pass)
		VALUES ($1, $2)
		RETURNING user_id
	`, userTable)

	row := tx.QueryRow(query, input.Email, input.Password)

	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value") {
			return id, models.UserAlreadyExist
		}
		return id, err
	}

	query = fmt.Sprintf(`
		INSERT INTO %s (user_id, name)
		VALUES ($1, $2)
	`, profileTable)

	_, err = tx.Query(query, id, input.Name)
	if err != nil {
		tx.Rollback()
		return id, err
	}

	tx.Commit()

	return id, nil
}

func (r *UserPostgres) AuthenticateUser(email, passHash string) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
		SELECT user_id, name, email FROM %s
		WHERE email=$1 AND hashed_pass=$2
	`, userTable)

	err := r.db.Get(&user, query, email, passHash)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return nil, models.EmailOrPassNotMatch
	}

	return &user, err
}

func (r *UserPostgres) GetUserById(userId int) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1
	`, userTable)

	err := r.db.Get(&user, query, userId)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return nil, models.UserDoesntExist
	}

	return &user, err
}

func (r *UserPostgres) ChangePassword(userId int, password string) error {

	query := fmt.Sprintf(`
		UPDATE %s
		SET hashed_pass = $1
		WHERE user_id = $2
	`, userTable)

	_, err := r.db.Exec(query, password, userId)

	return err
}
