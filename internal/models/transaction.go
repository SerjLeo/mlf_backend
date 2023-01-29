package models

type Transaction struct {
	Id              int        `json:"id" db:"transaction_id"`
	UserId          int        `json:"-" db:"user_id"`
	Amount          float64    `json:"amount" db:"amount"`
	Description     string     `json:"description" db:"description"`
	TransactionType bool       `json:"type" db:"type"`
	Currency        string     `json:"currency,omitempty" db:"currency"`
	CurrencyId      int        `json:"currency_id,omitempty" db:"currency_id"`
	CreatedAt       string     `json:"created_at" db:"created_at"`
	UpdatedAt       string     `json:"updated_at" db:"updated_at"`
	Categories      []Category `json:"categories,omitempty"`
}

type CreateTransactionInput struct {
	Amount          float64 `json:"amount" binding:"required" db:"amount"`
	Description     string  `json:"description" db:"description" binding:"max=255"`
	TransactionType bool    `json:"type" db:"type"`
	CategoriesIds   []int   `json:"categories,omitempty"`
	CurrencyId      int     `json:"currency_id" binding:"required" db:"currency_id"`
	AccountId       int     `json:"account_id" binding:"required" db:"currency_id"`
}
