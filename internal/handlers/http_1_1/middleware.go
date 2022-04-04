package http_1_1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
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
	stringId, exists := c.Get("userId")
	if !exists {
		return 0, errors.New("user id doesn't provided")
	}

	intId, ok := stringId.(int)
	if !ok {
		return 0, errors.New("wrong user id format")
	}

	return intId, nil
}
