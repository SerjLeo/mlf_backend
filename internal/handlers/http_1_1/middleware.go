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

func (h *RequestHandler) isUserAuthenticated(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "wrong auth header format")
		return
	}

	userId, err := h.services.User.CheckUserToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *RequestHandler) getUserId(c *gin.Context) (int, error) {
	stringId, exists := c.Get(userCtx)
	if !exists {
		return 0, errors.New("user id doesn't provided")
	}

	intId, ok := stringId.(int)
	if !ok {
		return 0, errors.New("wrong user id format")
	}

	return intId, nil
}

func (h *RequestHandler) withPagination(c *gin.Context) {
	paginationParams := models.PaginationParams{}

	err := c.BindQuery(&paginationParams)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Bad pagination params").Error())
		return
	}

	c.Set(paginationCtx, paginationParams)
}

func (h *RequestHandler) getPagination(c *gin.Context) (models.PaginationParams, error) {
	pagination, exists := c.Get(paginationCtx)
	if !exists {
		return models.PaginationParams{}, errors.New("pagination doesn't provided")
	}

	params, ok := pagination.(models.PaginationParams)
	if !ok {
		return models.PaginationParams{}, errors.New("Something wrong with pagination")
	}
	
	return params, nil
}
