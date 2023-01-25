package custom_errors

import "errors"

var (
	// global
	Unauthorized           = errors.New("credentials expired")
	ServerError            = errors.New("internal server error")
	PaginationParsingError = errors.New("error while parsing pagination")
	BadInput               = errors.New("error while parsing input")

	// auth
	UserAlreadyExist    = errors.New("user already exists")
	EmailOrPassNotMatch = errors.New("email or password doesn't match")
	UserDoesntExist     = errors.New("user doesn't exist")
	CurrencyNotFound    = errors.New("currency not found")

	// account
	AccountIDIsNotProvided = errors.New("account ID is not provided")
	AccountInvalidID       = errors.New("account ID is invalid")
	AccountNotExist        = errors.New("account doesn't exist")

	//category
	CategoryIdNotProvided = errors.New("category id is not provided")
	CategoryNotExist      = errors.New("category doesn't exist")
	CategoryInvalidID     = errors.New("category ID is invalid")

	//Transaction
	TransactionIdNotProvided = errors.New("transaction id is not provided")
)
