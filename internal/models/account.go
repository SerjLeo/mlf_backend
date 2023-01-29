package models

import "errors"

type Account struct {
	Id        int    `json:"id" db:"account_id"`
	UserId    int    `json:"-" db:"user_id"`
	Name      string `json:"name" db:"name"`
	Suspended bool   `json:"-" db:"suspended"`
	IsDefault bool   `json:"-" db:"is_default"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type AccountWithBalances struct {
	Id        int       `json:"id" db:"account_id"`
	UserId    int       `json:"-" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Suspended bool      `json:"-" db:"suspended"`
	IsDefault bool      `json:"-" db:"is_default"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
	Balances  []Balance `json:"balances"`
}

type CreateAccountInput struct {
	Name       string `json:"name" db:"name" binding:"required,min=3"`
	CurrencyId int    `json:"currency_id" db:"currency_id"`
}

func (i *CreateAccountInput) Validate() error {
	if len(i.Name) < 3 {
		return errors.New("name should be 3 digits or longer")
	}
	return nil
}

type UpdateAccountInput struct {
	Name      string `json:"name" db:"name" binding:"required,min=3"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func (i *UpdateAccountInput) Validate() error {
	if len(i.Name) < 3 {
		return errors.New("name should be 3 digits or longer")
	}
	return nil
}
