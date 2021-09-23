package models

type Category struct {
	CategoryId int    `json:"category_id" db:"category_id"'`
	UserId     int    `json:"user_id" db:"user_id"`
	Name       string `json:"name" db:"name"`
	Color      string `json:"color" db:"color"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

type CreateCategoryInput struct {
	Name  string `json:"name" db:"name" binding:"required"`
	Color string `json:"color" db:"color"`
}
