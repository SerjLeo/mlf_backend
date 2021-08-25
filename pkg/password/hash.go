package password

import (
	"crypto/sha1"
	"fmt"
)

type HashGenerator interface {
	EncodeString(source string) (string, error)
}

type SHA1Hash struct {
	salt string
}

func NewSHA1Hash(salt string) *SHA1Hash {
	return &SHA1Hash{salt: salt}
}

func (g *SHA1Hash) EncodeString(source string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(source)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(g.salt))), nil
}
