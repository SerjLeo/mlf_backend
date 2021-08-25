package models

type User struct {
	UserId   int    `json:"-" db:"user_id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
