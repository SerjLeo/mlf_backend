package http_1_1

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	paginationCtx       = "pagination"
)

func (h *HTTPRequestHandler) isUserAuthenticated(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, custom_errors.Unauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, custom_errors.Unauthorized)
		return
	}

	userId, err := h.services.CheckUserToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.Set(userCtx, userId)
}

func (h *HTTPRequestHandler) getUserId(ctx *gin.Context) (int, error) {
	stringId, exists := ctx.Get(userCtx)
	if !exists {
		return 0, custom_errors.Unauthorized
	}

	intId, ok := stringId.(int)
	if !ok {
		return 0, custom_errors.Unauthorized
	}

	return intId, nil
}

func (h *HTTPRequestHandler) withPagination(ctx *gin.Context) {
	paginationParams := models.PaginationParams{}

	err := ctx.BindQuery(&paginationParams)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.PaginationParsingError)
		return
	}

	ctx.Set(paginationCtx, paginationParams)
}

func (h *HTTPRequestHandler) getPagination(ctx *gin.Context) (models.PaginationParams, error) {
	pagination, exists := ctx.Get(paginationCtx)
	if !exists {
		return models.PaginationParams{}, nil
	}

	params, ok := pagination.(models.PaginationParams)
	if !ok {
		return models.PaginationParams{}, custom_errors.PaginationParsingError
	}

	return params, nil
}
