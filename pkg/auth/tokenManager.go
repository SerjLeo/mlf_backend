package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenManager interface {
	GenerateToken(userId int, ttl time.Duration) (string, error)
	ParseToken(token string) (*tokenClaims, error)
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

func (m *Manager) ParseToken(token string) (*tokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signKey), nil
	})

	if err != nil {
		return &tokenClaims{}, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return &tokenClaims{}, fmt.Errorf("error while getting user claims from token")
	}

	return claims, nil
}
