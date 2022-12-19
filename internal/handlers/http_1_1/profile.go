package http_1_1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *HTTPRequestHandler) initProfileRoutes(api *gin.RouterGroup) {
	auth := api.Group("/profile")
	{
		auth.GET("", h.getProfile)
		auth.POST("", h.editProfile)
	}
}

// @Summary Get user's profile
// @Tags profile
// @Description get existing user profile
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profile [get]
func (h *HTTPRequestHandler) getProfile(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	profile, err := h.services.GetUserProfile(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("Error getting profile:", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: profile,
	})
}

func (h *HTTPRequestHandler) editProfile(c *gin.Context) {

}
