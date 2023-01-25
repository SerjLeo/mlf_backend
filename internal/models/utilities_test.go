package models_test

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginationParams_ToString(t *testing.T) {
	testCases := []struct {
		name     string
		input    models.PaginationParams
		expected string
	}{
		{
			name: "only page",
			input: models.PaginationParams{
				Page: 3,
			},
			expected: "page=3",
		},
		{
			name: "only per page",
			input: models.PaginationParams{
				PerPage: 5,
			},
			expected: "perPage=5",
		},
		{
			name: "both params",
			input: models.PaginationParams{
				Page:    10,
				PerPage: 22,
			},
			expected: "page=10&perPage=22",
		},
		{
			name:     "no params",
			input:    models.PaginationParams{},
			expected: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.input.ToString(), tt.expected)
		})
	}
}
