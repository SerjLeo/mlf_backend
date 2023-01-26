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

func TestGetUserCategories(t *testing.T) {
	router, mock := SetupTest(t)

	categories := []models.Category{
		{
			Id: 1,
		},
	}
	correctPagination := models.PaginationParams{
		Page: 1,
	}
	missingPagination := models.PaginationParams{
		Page: 2,
	}

	mock.On("GetUserCategories", validUserId, correctPagination).Return(categories, nil)
	mock.On("GetUserCategories", validUserId, missingPagination).Return([]models.Category{}, nil)
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
			expectedOutput: dataResponse{Data: categories},
		},
		{
			name:       "with missing pagination",
			pagination: missingPagination,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: []models.Category{}},
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
			w := PerformRequest(router, "GET", fmt.Sprintf("/api/category%s?%s", http_1_1.GetCategoriesRoute, tt.pagination.ToString()), "", tt.headers...)

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

func TestGetUserCategoryById(t *testing.T) {
	router, mock := SetupTest(t)

	category := &models.Category{}

	invalidCatID := "sometext"

	mock.On("GetUserCategoryById", validUserId, existingIdParamValue).Return(category, nil)
	mock.On("GetUserCategoryById", validUserId, notExistingIdParamValue).Return(nil, custom_errors.CategoryNotExist)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		catIdParam     string
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name:       "with valid existing category id",
			catIdParam: existingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: category},
		},
		{
			name:       "with valid not existing account id",
			catIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.CategoryNotExist.Error()},
		},
		{
			name:       "with not valid category id",
			catIdParam: invalidCatID,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.CategoryInvalidID.Error()},
		},
		{
			name:           "unauthorized user",
			catIdParam:     existingIdParam,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "GET", fmt.Sprintf("/api/category/%s", tt.catIdParam), "", tt.headers...)

			CheckResults(t, w, tt.expectedCode, tt.expectedOutput)
		})
	}
}

func TestCreateCategory(t *testing.T) {
	router, mock := SetupTest(t)

	invalidInput := models.CreateCategoryInput{
		Name:  "",
		Color: "#ffffff",
	}

	validInput := models.CreateCategoryInput{
		Name:  "new name",
		Color: "#ffffff",
	}

	category := &models.Category{}

	mock.On("CreateCategory", validUserId, &validInput).Return(category, nil)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		expectedCode   int
		payload        models.CreateCategoryInput
		expectedOutput dataResponse
	}{
		{
			name:    "with valid payload",
			payload: validInput,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: category},
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
			w := PerformRequest(router, "POST", "/api/category", tt.payload, tt.headers...)

			CheckResults(t, w, tt.expectedCode, tt.expectedOutput)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	router, mock := SetupTest(t)

	invalidInput := models.UpdateCategoryInput{
		Name: "",
	}

	validInput := models.UpdateCategoryInput{
		Name: "new name",
	}

	category := &models.Category{}

	validCatId := 1
	invalidCatId := 2

	mock.On("UpdateCategory", validUserId, validCatId, &validInput).Return(category, nil)
	mock.On("UpdateCategory", validUserId, invalidCatId, &validInput).Return(nil, custom_errors.CategoryNotExist)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		accIdParam     string
		expectedCode   int
		payload        models.UpdateCategoryInput
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
			expectedOutput: dataResponse{Data: category},
		},
		{
			name:       "with valid payload and not existing id",
			payload:    validInput,
			accIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.CategoryNotExist.Error()},
		},
		{
			name:       "with invalid id",
			payload:    validInput,
			accIdParam: invalidIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.CategoryInvalidID.Error()},
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
			w := PerformRequest(router, "PUT", fmt.Sprintf("/api/category/%s", tt.accIdParam), tt.payload, tt.headers...)

			CheckResults(t, w, tt.expectedCode, tt.expectedOutput)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	router, mock := SetupTest(t)

	existingCatId := 1
	notExistingCatId := 2
	invalidCatID := "sometext"

	mock.On("DeleteCategory", validUserId, existingCatId).Return(nil)
	mock.On("DeleteCategory", validUserId, notExistingCatId).Return(custom_errors.CategoryNotExist)
	mock.On("CheckUserToken", validToken).Return(validUserId, nil)

	testCases := []struct {
		name           string
		headers        []header
		catIdParam     string
		expectedCode   int
		expectedOutput dataResponse
	}{
		{
			name:       "with valid existing category id",
			catIdParam: existingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusOK,
			expectedOutput: dataResponse{Data: true},
		},
		{
			name:       "with valid not existing category id",
			catIdParam: notExistingIdParam,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusInternalServerError,
			expectedOutput: dataResponse{Error: custom_errors.CategoryNotExist.Error()},
		},
		{
			name:       "with not valid category id",
			catIdParam: invalidCatID,
			headers: []header{
				{Key: "Authorization", Value: validTokenHeader},
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: dataResponse{Error: custom_errors.CategoryInvalidID.Error()},
		},
		{
			name:           "unauthorized user",
			catIdParam:     existingIdParam,
			headers:        []header{},
			expectedCode:   http.StatusUnauthorized,
			expectedOutput: dataResponse{Error: custom_errors.Unauthorized.Error()},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "DELETE", fmt.Sprintf("/api/category/%s", tt.catIdParam), "", tt.headers...)

			CheckResults(t, w, tt.expectedCode, tt.expectedOutput)
		})
	}
}
