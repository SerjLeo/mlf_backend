package models

type Currency struct {
	CurrencyId int    `json:"currency_id" db:"currency_id"`
	Name       string `json:"name" db:"name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
