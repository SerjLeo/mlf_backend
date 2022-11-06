package models

type Currency struct {
	CurrencyId int `json:"currency_id"`
	Name       string
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
