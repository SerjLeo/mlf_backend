package models

type Transaction struct {
	TransactionId   int        `json:"transaction_id" db:"transaction_id"`
	UserId          int        `json:"user_id" db:"user_id"`
	Amount          float64    `json:"amount" binding:"required" db:"amount"`
	Description     string     `json:"description" db:"description" binding:"max=255"`
	TransactionType bool       `json:"type" db:"type"`
	CreatedAt       string     `json:"created_at" db:"created_at"`
	UpdatedAt       string     `json:"updated_at" db:"updated_at"`
	Categories      []Category `json:"categories,omitempty"`
}

type CreateTransactionInput struct {
	Amount          float64 `json:"amount" binding:"required" db:"amount"`
	Description     string  `json:"description" db:"description" binding:"max=255"`
	TransactionType bool    `json:"type" db:"type"`
	CategoriesIds   []int   `json:"categories,omitempty"`
}
