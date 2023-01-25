package http_1_1_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSignInHandler(t *testing.T) {
	router, mock := SetupTest(t)

	mock.On("SignIn", "correctEmail", "correctPassword").Return(validToken, nil)
	mock.On("SignIn", "incorrectEmail", "incorrectPassword").Return("", errors.New("user dont exist"))

	t.Run("authorization succeeded", func(t *testing.T) {
		payload := http_1_1.SignInInput{Email: "correctEmail", Password: "correctPassword"}
		w := PerformRequest(router, "POST", fmt.Sprintf("/api/auth%s", http_1_1.SignInRoute), payload)

		assert.Equal(t, w.Code, http.StatusOK)
		mock.AssertCalled(t, "SignIn", payload.Email, payload.Password)
		result := dataResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Error("Fail to parse response")
		}
		assert.ObjectsAreEqual(result, dataResponse{Data: validToken})
	})

	t.Run("wrong credentials", func(t *testing.T) {
		payload := http_1_1.SignInInput{Email: "incorrectEmail", Password: "incorrectPassword"}
		w := PerformRequest(router, "POST", fmt.Sprintf("/api/auth%s", http_1_1.SignInRoute), payload)

		assert.Equal(t, w.Code, http.StatusBadRequest)
		mock.AssertCalled(t, "SignIn", payload.Email, payload.Password)
	})

	invalidDataTestCases := []struct {
		name           string
		payload        http_1_1.SignInInput
		expectedStatus int
	}{
		{
			name:           "no email provided",
			payload:        http_1_1.SignInInput{Email: "", Password: "incorrectPassword"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no password provided",
			payload:        http_1_1.SignInInput{Email: "incorrectEmail", Password: ""},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no email and password provided",
			payload:        http_1_1.SignInInput{Email: "", Password: ""},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range invalidDataTestCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "POST", fmt.Sprintf("/api/auth%s", http_1_1.SignInRoute), tt.payload)

			assert.Equal(t, w.Code, tt.expectedStatus)
			mock.AssertNotCalled(t, "SignIn", tt.payload.Email, tt.payload.Password)
		})
	}

}

func TestSignUpHandler(t *testing.T) {
	router, mock := SetupTest(t)

	mockUser := models.CreateUserInput{
		Name:     "username",
		Email:    "example@mail.com",
		Password: "password",
	}
	mock.On("Create", &mockUser).Return(validToken, nil)

	t.Run("sing up succeeded", func(t *testing.T) {
		payload := mockUser
		w := PerformRequest(router, "POST", fmt.Sprintf("/api/auth%s", http_1_1.SignUpRoute), payload)

		assert.Equal(t, w.Code, http.StatusCreated)
		mock.AssertCalled(t, "Create", &mockUser)
		result := dataResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Error("Fail to parse response")
		}
		assert.ObjectsAreEqual(result, dataResponse{Data: validToken})
	})

	invalidDataTestCases := []struct {
		name           string
		payload        models.CreateUserInput
		expectedStatus int
	}{
		{
			name:           "no email provided",
			payload:        models.CreateUserInput{Email: "", Password: "password", Name: "username"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no password provided",
			payload:        models.CreateUserInput{Email: "example@mail.com", Password: "", Name: "username"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "no name provided",
			payload:        models.CreateUserInput{Email: "example@mail.com", Name: "", Password: "password"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "email is not valid",
			payload:        models.CreateUserInput{Email: "examplemailcom", Password: "password", Name: "username"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "password length less than 6",
			payload:        models.CreateUserInput{Email: "example@mail.com", Password: "pass", Name: "username"},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range invalidDataTestCases {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(router, "POST", fmt.Sprintf("/api/auth%s", http_1_1.SignUpRoute), tt.payload)

			assert.Equal(t, w.Code, tt.expectedStatus)
			mock.AssertNotCalled(t, "Create", &models.CreateUserInput{
				Email:    tt.payload.Email,
				Name:     tt.payload.Name,
				Password: tt.payload.Password,
			})
		})
	}
}
