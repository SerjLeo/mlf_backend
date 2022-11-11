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

func (r *UserPostgres) Create(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (name, email, hashed_pass)
		VALUES ($1, $2, $3)
		RETURNING user_id
	`, userTable)

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
	query := fmt.Sprintf(`
		SELECT user_id, name, email FROM %s
		WHERE email=$1 AND hashed_pass=$2
	`, userTable)

	err := r.db.Get(&user, query, email, passHash)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return user, models.EmailOrPassNotMatch
	}

	return user, err
}

func (r *UserPostgres) GetUserById(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
		SELECT email, %s.name, %s.name as currency FROM %s INNER JOIN %s ON %s.currency_id=%s.currency_id
		WHERE user_id=$1
	`, userTable, currencyTable, userTable, currencyTable, userTable, currencyTable)

	err := r.db.Get(&user, query, userId)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return user, models.UserDoesntExist
	}

	return user, err
}

func (r *UserPostgres) GetFullUserById(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id=$1
	`, userTable)

	err := r.db.Get(&user, query, userId)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		return user, models.UserDoesntExist
	}

	return user, err
}

func (r *UserPostgres) UpdateUser(userId int, input models.User) error {

	fmt.Printf("%+v", input)

	query := fmt.Sprintf(`
		UPDATE %s
		SET name = $1, currency_id = $2, updated_at = $3
		WHERE user_id = $4
		RETURNING *
	`, userTable)

	_, err := r.db.Exec(query, input.Name, input.CurrencyId, input.UpdatedAt, userId)

	return err
}
