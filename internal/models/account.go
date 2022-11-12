package models

type Account struct {
	Id        int    `json:"id" db:"account_id"`
	Name      string `json:"name" db:"name"`
	Suspended bool   `json:"-" db:"suspended"`
	IsDefault bool   `json:"-" db:"is_default"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
