package models

type User struct {
	UserId      int    `json:"-" db:"user_id"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required" db:"hashed_pass"`
	IsConfirmed bool   `json:"is_confirmed" db:"is_confirmed"`
	Currency    string `json:"currency,omitempty" db:"currency"`
	CurrencyId  int    `json:"currency_id,omitempty" db:"currency_id"`
	UserRole    int    `json:"user_role" db:"user_role"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
