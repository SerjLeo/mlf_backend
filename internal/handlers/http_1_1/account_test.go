package http_1_1_test

import (
	"encoding/json"
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetUserAccounts(t *testing.T) {
	router, mock := SetupTest(t)

	accounts := []models.AccountWithBalances{
		{
			Id:        0,
			Name:      "",
			Suspended: false,
			IsDefault: false,
			CreatedAt: "",
			UpdatedAt: "",
			Balances:  nil,
		},
	}
	correctPagination := models.PaginationParams{
		Page: 1,
	}
	missingPagination := models.PaginationParams{
		Page: 2,
	}

	mock.On("GetAccounts", correctPagination, validUserId).Return(accounts, nil)
	mock.On("GetAccounts", missingPagination, validUserId).Return([]models.AccountWithBalances{}, nil)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		pagination     models.PaginationParams
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name:       "with correct pagination",
			pagination: correctPagination,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: accounts},
		},
		{
			name:       "with missing pagination",
			pagination: missingPagination,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: []models.AccountWithBalances{}},
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
			w := PerformRequest(router, "GET", fmt.Sprintf("/api/account%s?%s", http_1_1.GetAccountsRoute, tt.pagination.ToString()), "", tt.headers...)

			result := dataResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Error("Fail to parse response")
			}

			assert.Equal(t, tt.expectedCode, w.Code)
			if result.Data != nil {
				assert.ObjectsAreEqual(tt.expectedOutput.Data, result.Data)
			}

			if result.Error != "" {
				assert.Truef(
					t,
					strings.Contains(result.Error,
						tt.expectedOutput.Error),
					"expected error message \n \"%s\" \n inclusing \"%s\"",
					result.Error,
					tt.expectedOutput.Error,
				)
			}
		})
	}

}

func TestGetUserAccountById(t *testing.T) {
	router, mock := SetupTest(t)

	account := &models.AccountWithBalances{}

	invalidAccID := "sometext"
	existingId := 1
	notExistingId := 2

	mock.On("GetAccountById", existingId, validUserId).Return(account, nil)
	mock.On("GetAccountById", notExistingId, validUserId).Return(nil, custom_errors.AccountNotExist)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		accIdParam     string
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name:       "with valid existing account id",
			accIdParam: existingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: account},
		},
		{
			name:       "with valid not existing account id",
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.AccountNotExist.Error()},
		},
		{
			name:       "with not valid account id",
			accIdParam: invalidAccID,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.AccountInvalidID.Error()},
		},
		{
			name:           "unauthorized user",
			accIdParam:     existingIdParam,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "GET", fmt.Sprintf("/api/account/%s", tt.accIdParam), "", tt.headers...)

			result := dataResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Error("Fail to parse response")
			}

			assert.Equal(t, tt.expectedCode, w.Code)

			if result.Data != nil {
				assert.ObjectsAreEqual(tt.expectedOutput.Data, result.Data)
			}

			if result.Error != "" {
				assert.Truef(
					t,
					strings.Contains(result.Error,
						tt.expectedOutput.Error),
					"expected error message \n \"%s\" \n inclusing \"%s\"",
					result.Error,
					tt.expectedOutput.Error,
				)
			}
		})
	}

}

func TestCreateUserAccount(t *testing.T) {
	router, mock := SetupTest(t)

	invalidInput := models.CreateAccountInput{
		Name: "",
	}

	validInput := models.CreateAccountInput{
		Name:       "account",
		CurrencyId: 1,
	}

	account := &models.AccountWithBalances{}

	mock.On("CreateAccount", &validInput, validUserId).Return(account, nil)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		expectedCode   int
		payload        models.CreateAccountInput
		expectedOutput dataResponse
	}{
		{
			name:    "with valid payload",
			payload: validInput,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusCreated,
			expectedOutput: dataResponse{Data: account},
		},
		{
			name:    "with invalid payload",
			payload: invalidInput,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.BadInput.Error()},
		},
		{
			name:           "unauthorized user",
			payload:        validInput,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "POST", "/api/account", tt.payload, tt.headers...)

			fmt.Printf("%+v", w.Body)

			result := dataResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Error("Fail to parse response")
			}

			assert.Equal(t, tt.expectedCode, w.Code)

			if result.Data != nil {
				assert.ObjectsAreEqual(tt.expectedOutput.Data, result.Data)
			}

			if result.Error != "" {
				assert.Truef(
					t,
					strings.Contains(result.Error,
						tt.expectedOutput.Error),
					"expected error message \n \"%s\" \n inclusing \"%s\"",
					result.Error,
					tt.expectedOutput.Error,
				)
			}
		})
	}

}

func TestUpdateUserAccount(t *testing.T) {
	router, mock := SetupTest(t)

	invalidInput := models.UpdateAccountInput{
		Name: "",
	}

	validInput := models.UpdateAccountInput{
		Name: "new name",
	}

	account := &models.AccountWithBalances{}

	validAccId := 1
	invalidAccId := 2

	mock.On("UpdateAccount", validAccId, validUserId, &validInput).Return(account, nil)
	mock.On("UpdateAccount", invalidAccId, validUserId, &validInput).Return(nil, custom_errors.AccountNotExist)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		accIdParam     string
		expectedCode   int
		payload        models.UpdateAccountInput
		expectedOutput dataResponse
	}{
		{
			name:       "with valid payload and existing id",
			payload:    validInput,
			accIdParam: existingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: account},
		},
		{
			name:       "with valid payload and not existing id",
			payload:    validInput,
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.AccountNotExist.Error()},
		},
		{
			name:       "with invalid id",
			payload:    validInput,
			accIdParam: invalidIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.AccountInvalidID.Error()},
		},
		{
			name:       "with invalid payload",
			payload:    invalidInput,
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.BadInput.Error()},
		},
		{
			name:       "with invalid account Id",
			payload:    invalidInput,
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.BadInput.Error()},
		},
		{
			name:           "unauthorized user",
			payload:        validInput,
			accIdParam:     existingIdParam,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "PUT", fmt.Sprintf("/api/account/%s", tt.accIdParam), tt.payload, tt.headers...)

			result := dataResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Error("Fail to parse response")
			}

			assert.Equal(t, tt.expectedCode, w.Code)

			if result.Data != nil {
				assert.ObjectsAreEqual(tt.expectedOutput.Data, result.Data)
			}

			if result.Error != "" {
				assert.Truef(
					t,
					strings.Contains(result.Error,
						tt.expectedOutput.Error),
					"expected error message \n \"%s\" \n inclusing \"%s\"",
					result.Error,
					tt.expectedOutput.Error,
				)
			}
		})
	}

}

func TestDeleteUserAccount(t *testing.T) {
	router, mock := SetupTest(t)

	invalidAccID := "sometext"
	existingAccId := 1
	notExistingAccId := 2

	mock.On("SoftDeleteAccount", existingAccId, validUserId).Return(nil)
	mock.On("SoftDeleteAccount", notExistingAccId, validUserId).Return(custom_errors.AccountNotExist)
	mock.On("CheckUserToken", validToken).Return(1, nil)

	testCases := []struct {
		name           string
		headers        []header
		accIdParam     string
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name:       "with valid existing account id",
			accIdParam: existingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: true},
		},
		{
			name:       "with valid not existing account id",
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.AccountNotExist.Error()},
		},
		{
			name:       "with not valid account id",
			accIdParam: invalidAccID,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.AccountInvalidID.Error()},
		},
		{
			name:           "unauthorized user",
			accIdParam:     existingIdParam,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "DELETE", fmt.Sprintf("/api/account/%s", tt.accIdParam), "", tt.headers...)

			result := dataResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Error("Fail to parse response")
			}

			assert.Equal(t, tt.expectedCode, w.Code)

			if result.Data != nil {
				assert.ObjectsAreEqual(tt.expectedOutput.Data, result.Data)
			}

			if result.Error != "" {
				assert.Truef(
					t,
					strings.Contains(result.Error,
						tt.expectedOutput.Error),
					"expected error message \n \"%s\" \n inclusing \"%s\"",
					result.Error,
					tt.expectedOutput.Error,
				)
			}
		})
	}

}
