package auth

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewTokenManager(t *testing.T) {

	t.Run("should return error if sign key is not provided", func(t *testing.T) {
		_, err := NewTokenManager("")
		assert.Error(t, err)
	})

	t.Run("should return manager and nil error if sign key is provided", func(t *testing.T) {
		manager, err := NewTokenManager("somekey")
		assert.Nil(t, err)
		assert.NotEmpty(t, manager)
	})

}
