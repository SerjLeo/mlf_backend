package http_1_1

import "github.com/gin-gonic/gin"

type dataResponse struct {
	Data interface{}
}

type errorResponse struct {
	Error string
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
