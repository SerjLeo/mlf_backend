package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	EditProfileRoute = "/profile"
)

func (h *HTTPRequestHandler) initUserRoutes(api *gin.RouterGroup) {
	auth := api.Group("/user")
	{
		auth.POST(EditProfileRoute, h.userEditProfile)
	}
}

// @Summary Edit profile
// @Tags user
// @Description Makes changes in user profile
// @Accept  json
// @Produce  json
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/profile [post]
func (h *HTTPRequestHandler) userEditProfile(c *gin.Context) {
	h.isUserAuthenticated(c)

	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while checking authorization")
		return
	}

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.UpdateUserProfile(userId, input)
	if err != nil {
		fmt.Println(err.Error())
		newErrorResponse(c, http.StatusNotFound, "user profile doesn't exists")
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: userResponse{
			Name:     user.Name,
			Email:    user.Email,
			Currency: user.Currency,
		},
	})
}
