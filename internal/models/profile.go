package models

type Profile struct {
	Id           int    `json:"-" db:"profile_id"`
	Name         string `json:"name" db:"name"`
	CurrencyId   int    `json:"currency_id" db:"currency_id"`
	CurrencyName string `json:"currency" db:"currency"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

type FullProfile struct {
	Id           int    `json:"-" db:"profile_id"`
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"'`
	CurrencyId   int    `json:"currency_id" db:"currency_id"`
	CurrencyName string `json:"currency" db:"currency"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

type UpdateProfileInput struct {
	Name       string `json:"name,omitempty" db:"name"`
	CurrencyId int    `json:"currency_id,omitempty" db:"currency_id"`
}
