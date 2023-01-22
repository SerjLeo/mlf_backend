package http_1_1_test

import (
	"bytes"
	"encoding/json"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/mocks"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type header struct {
	Key   string
	Value string
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

func SetupTest(t *testing.T) (*gin.Engine, *mocks.Service) {
	service := mocks.NewService(t)
	handler := http_1_1.NewRequestHandler(service, "")
	gin.SetMode(gin.TestMode)
	return handler.InitRoutes(), service
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
