package models

type Profile struct {
	Id           int    `json:"-" db:"profile_id"`
	Name         string `json:"name" db:"name"`
	CurrencyId   int    `json:"currency_id" db:"currency_id"`
	CurrencyName string `json:"currency" db:"currency"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}
