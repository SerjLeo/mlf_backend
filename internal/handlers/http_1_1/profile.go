package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *HTTPRequestHandler) initProfileRoutes(api *gin.RouterGroup) {
	profile := api.Group("/profile", h.isUserAuthenticated)
	{
		profile.GET("", h.getProfile)
		profile.PUT("", h.editProfile)
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
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	profile, err := h.services.GetUserProfile(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting profile"))
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
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	input := &models.UpdateProfileInput{}
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	fmt.Printf("%+v \n", input)

	profile, err := h.services.UpdateProfile(input, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: profile,
	})
}
