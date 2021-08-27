package http_1_1

import "github.com/gin-gonic/gin"

type dataResponse struct {
	data interface{}
}

type dataWithPaginationResponse struct {
	data interface{}
	meta metaParams
}

type metaParams struct {
	page    int
	perPage int
	total   int
}

type errorResponse struct {
	Error string
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
