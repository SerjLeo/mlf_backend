package models

import "errors"

var (
	UserAlreadyExist = errors.New("user already exists")
	EmailOrPassNotMatch = errors.New("email or password doesn't match")
)
