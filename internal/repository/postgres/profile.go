package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) CreateProfile(userId int, name string) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, name)
		VALUES ($1, $2)
	`, profileTable)

	row := r.db.QueryRow(query, userId, name)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *ProfilePostgres) UpdateProfile(userId int) error {
	return nil
}

func (r *ProfilePostgres) GetUserProfile(userId int) (*models.FullProfile, error) {
	profile := models.FullProfile{}
	query := fmt.Sprintf(`
		SELECT profile_id, name, email, currency_id, %s.name as currency, created_at, updated_at
		FROM %s INNER JOIN %s ON %s.currency_id=%s.currency_id INNER JOIN %s ON %s.user_id=%s.user_id
		WHERE user_id=$1
	`, currencyTable, profileTable, currencyTable, profileTable, currencyTable, userTable, profileTable, userTable)
	err := r.db.Get(&profile, query, userId)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfilePostgres) DeleteProfile(userId, profileId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE profileId=$1 AND userId=$2
	`, profileTable)

	_, err := r.db.Exec(query, profileId, userId)
	return err
}
