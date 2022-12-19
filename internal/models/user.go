package models

type User struct {
	Id           int    `json:"-" db:"user_id"`
	Email        string `json:"email" binding:"required,email"`
	PasswordHash string `json:"-" binding:"required" db:"hashed_pass"`
	IsConfirmed  bool   `json:"-" db:"is_confirmed"`
	UserRole     int    `json:"-" db:"user_role"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
}
