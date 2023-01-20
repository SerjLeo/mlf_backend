package models

type Currency struct {
	CurrencyId int    `json:"id" db:"currency_id"`
	Name       string `json:"name" db:"name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

type CurrencyType int8

const (
	USD CurrencyType = 1
	EUR              = 2
	TL               = 3
	RUB              = 4
)
