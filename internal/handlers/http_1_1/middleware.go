package http_1_1

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong auth header format")
		return
	}

	userId, err := h.services.CheckUserToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func (h *HTTPRequestHandler) getUserId(ctx *gin.Context) (int, error) {
	stringId, exists := ctx.Get(userCtx)
	if !exists {
		return 0, errors.New("user id doesn't provided")
	}

	intId, ok := stringId.(int)
	if !ok {
		return 0, errors.New("wrong user id format")
	}

	return intId, nil
}

func (h *HTTPRequestHandler) withPagination(ctx *gin.Context) {
	paginationParams := models.PaginationParams{}

	err := ctx.BindQuery(&paginationParams)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, errors.Wrap(err, "Bad pagination params").Error())
		return
	}

	ctx.Set(paginationCtx, paginationParams)
}

func (h *HTTPRequestHandler) getPagination(ctx *gin.Context) (models.PaginationParams, error) {
	pagination, exists := ctx.Get(paginationCtx)
	if !exists {
		return models.PaginationParams{}, errors.New("pagination doesn't provided")
	}

	params, ok := pagination.(models.PaginationParams)
	if !ok {
		return models.PaginationParams{}, errors.New("Something wrong with pagination")
	}

	return params, nil
}
