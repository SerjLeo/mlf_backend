package email

import (
	"errors"
	"regexp"
)

const (
	minEmailLen = 3
	maxEmailLen = 255
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type SendInput struct {
	To      string
	Body    string
	Subject string
}

func IsEmailValid(email string) bool {
	if len(email) < minEmailLen || len(email) > maxEmailLen {
		return false
	}

	return emailRegex.MatchString(email)
}

func (i *SendInput) Validate() error {
	if !IsEmailValid(i.To) || i.To == "" {
		return errors.New("invalid or empty target email")
	}

	if i.Subject == "" || i.Body == "" {
		return errors.New("empty subject or body")
	}

	return nil
}
