package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	GenerateToken(userId int, ttl time.Duration) (string, error)
	ParseToken(token string) (string, error)
}

type Manager struct {
	signKey string
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewTokenManager(signKey string) (*Manager, error) {
	if signKey == "" {
		return nil, errors.New("empty signKey for token manager")
	}
	return &Manager{signKey: signKey}, nil
}

func (m *Manager) GenerateToken(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(m.signKey))
}

func (m *Manager) ParseToken(token string) (string, error) {
	return "", nil
}
