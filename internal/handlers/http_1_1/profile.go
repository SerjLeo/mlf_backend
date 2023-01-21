package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *HTTPRequestHandler) initProfileRoutes(api *gin.RouterGroup) {
	auth := api.Group("/profile", h.isUserAuthenticated)
	{
		auth.GET("", h.getProfile)
		auth.PUT("", h.editProfile)
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
func (h *HTTPRequestHandler) getProfile(ctx *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	profile, err := h.services.GetUserProfile(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error getting profile:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: profile,
	})
}

// @Summary Update user's profile
// @Tags profile
// @Description update profile fields
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profile [put]
func (h *HTTPRequestHandler) editProfile(ctx *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	input := &models.UpdateProfileInput{}
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("%+v \n", input)

	profile, err := h.services.UpdateProfile(input, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: profile,
	})
}
