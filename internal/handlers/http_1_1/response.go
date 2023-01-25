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
	Error string `json:"error"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, errorResponse{err.Error()})
}
