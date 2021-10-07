package models

type Transaction struct {
	transactionId   int        `json:"transaction_id" db:"transaction_id"`
	userId          int        `json:"user_id" db:"user_id"`
	amount          float64    `json:"amount" binding:"required" db:"amount"`
	description     string     `json:"description" db:"description"`
	transactionType bool       `json:"type" binding:"required" db:"type"`
	createdAt       string     `json:"created_at" db:"created_at"`
	updatedAt       string     `json:"updated_at" db:"created_at"`
	categories      []Category `json:"categories,omitempty"`
}
