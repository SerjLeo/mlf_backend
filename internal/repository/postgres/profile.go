package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *ProfilePostgres) UpdateProfile(input *models.UpdateProfileInput, userId int) error {
	query := fmt.Sprintf(`UPDATE %s SET `, profileTable)
	qParts := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	counter := 1

	if input.Name != "" {
		qParts = append(qParts, fmt.Sprintf("name=$%d", counter))
		args = append(args, input.Name)
		counter++
	}

	if input.CurrencyId != 0 {
		qParts = append(qParts, fmt.Sprintf("currency_id=$%d", counter))
		args = append(args, input.CurrencyId)
		counter++
	}

	args = append(args, userId)

	query = query + strings.Join(qParts, ",") + fmt.Sprintf(" WHERE user_id=$%d", counter)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ProfilePostgres) GetUserProfile(userId int) (*models.FullProfile, error) {
	profile := models.FullProfile{}
	query := fmt.Sprintf(`
		SELECT prof.profile_id, prof.name, email, cur.currency_id, cur.name as currency, prof.created_at, prof.updated_at
		FROM %s AS prof
		INNER JOIN %s AS cur ON prof.currency_id=cur.currency_id
		INNER JOIN %s AS u ON u.user_id=prof.user_id
		WHERE prof.user_id=$1
	`, profileTable, currencyTable, userTable)

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
