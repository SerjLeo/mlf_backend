package models

type Balance struct {
	Id           int     `json:"id" db:"balance_id"`
	Amount       float64 `json:"amount" db:"amount"`
	CurrencyId   int     `json:"currency_id" db:"currency_id"`
	CurrencyName string  `json:"currency" db:"currency"`
	UserId       int     `json:"-" db:"user_id"`
	CreatedAt    string  `json:"created_at" db:"created_at"`
	UpdatedAt    string  `json:"updated_at" db:"updated_at"`
}

type BalanceOfCurrency struct {
	Total        float64 `json:"total" db:"total"`
	CurrencyId   int     `json:"currency_id" db:"currency_id"`
	CurrencyName string  `json:"currency" db:"currency"`
	UserId       int     `json:"-" db:"user_id"`
}
