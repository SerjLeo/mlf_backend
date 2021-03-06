package http_1_1

import "github.com/gin-gonic/gin"

type dataResponse struct {
	Data interface{} `json:"data"`
}

type dataWithPaginationResponse struct {
	Data interface{} `json:"data"`
	Meta metaParams  `json:"meta"`
}

type metaParams struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
}

type errorResponse struct {
	Error string
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
