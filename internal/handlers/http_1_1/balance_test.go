package http_1_1_test

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"net/http"
	"testing"
)

func TestGetTotalBalances(t *testing.T) {
	router, mock := SetupTest(t)

	balances := []models.BalanceOfCurrency{
		{},
	}

	mock.On("GetUserBalancesAmount", validUserId).Return(balances, nil)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name: "with correct data",
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: balances},
		},
		{
			name:           "unauthorized user",
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "GET", fmt.Sprintf("/api/balance%s", http_1_1.GetTotalBalancesRoute), "", tt.headers...)

			CheckResults(t, w, tt.expectedCode, tt.expectedOutput)
		})
	}

}
