package models

import "errors"

var (
	UserAlreadyExist    = errors.New("user already exists")
	EmailOrPassNotMatch = errors.New("email or password doesn't match")
	UserDoesntExist     = errors.New("user doesnt exist")
	CurrencyNotFound    = errors.New("currency not found")
)
