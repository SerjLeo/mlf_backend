package models

type Category struct {
	Id        int    `json:"id" db:"category_id"'`
	UserId    int    `json:"-" db:"user_id"`
	Name      string `json:"name" db:"name" binding:"max=255"`
	Color     string `json:"color" db:"color"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type CreateCategoryInput struct {
	Name  string `json:"name" db:"name" binding:"required" binding:"max=255"`
	Color string `json:"color" db:"color"`
}
