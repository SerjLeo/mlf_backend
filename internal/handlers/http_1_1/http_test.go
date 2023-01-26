package http_1_1_test

import (
	"bytes"
	"encoding/json"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	validUserId             = 1
	existingIdParam         = "1"
	notExistingIdParam      = "2"
	existingIdParamValue    = 1
	notExistingIdParamValue = 2
	invalidIdParam          = "text"
	validToken              = "token"
	validTokenHeader        = "Bearer token"
)

type header struct {
	Key   string
	Value string
}

type dataResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SetupTest(t *testing.T) (*gin.Engine, *mocks.Service) {
	service := mocks.NewService(t)
	handler := http_1_1.NewRequestHandler(service, "")
	gin.SetMode(gin.TestMode)
	return handler.InitRoutes(), service
}

func CheckResults(t *testing.T, w *httptest.ResponseRecorder, expectedCode int, expectedOutput dataResponse) {
	result := dataResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Error("Fail to parse response")
	}

	assert.Equal(t, expectedCode, w.Code)
	if result.Data != nil {
		assert.ObjectsAreEqual(expectedOutput.Data, result.Data)
	}

	if result.Error != "" {
		assert.Truef(
			t,
			strings.Contains(result.Error,
				expectedOutput.Error),
			"expected error message \n \"%s\" \n inclusing \"%s\"",
			result.Error,
			expectedOutput.Error,
		)
	}
}

func PerformRequest(r http.Handler, method, path string, body interface{}, headers ...header) *httptest.ResponseRecorder {
	encodedBody, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	bodyReader := bytes.NewReader(encodedBody)
	req := httptest.NewRequest(method, path, bodyReader)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	//req.Write()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
