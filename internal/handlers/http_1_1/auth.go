package http_1_1

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	CheckTokenRoute      = "/check"
	SignInRoute          = "/sign-in"
	SignUpRoute          = "/sign-up"
	SignUpWithEmailRoute = "/sign-up-with-email"
	RestorePasswordRoute = "/restore-password"
)

func (h *HTTPRequestHandler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.GET(CheckTokenRoute, h.userCheckToken)
		auth.POST(SignInRoute, h.userSignIn)
		auth.POST(SignUpRoute, h.userSignUp)
		auth.POST(SignUpWithEmailRoute, h.userSignUpWithEmail)
		auth.POST(RestorePasswordRoute, h.userRestorePassword)
	}
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary User sign-in with email and password
// @Tags auth
// @Description returns auth JWT
// @Accept  json
// @Produce  json
// @Param payload body SignInInput true "info for user's login"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *HTTPRequestHandler) userSignIn(ctx *gin.Context) {
	var input SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	token, err := h.services.SignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{Data: token})
}

type SignUpInput struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// @Summary User sign-up with name, email and password
// @Tags auth
// @Description creates user and returns auth JWT
// @Accept  json
// @Produce  json
// @Param payload body SignUpInput true "data for user creation"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *HTTPRequestHandler) userSignUp(ctx *gin.Context) {
	var input SignUpInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	token, err := h.services.Create(&models.CreateUserInput{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	})

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, dataResponse{Data: token})
}

type signUpWithEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

// @Summary User sign-up with email only
// @Tags auth
// @Description creates user with email, password generates automatically and returns auth JWT
// @Accept  json
// @Produce  json
// @Param payload body signUpWithEmailInput true "email for user creation"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up-with-email [post]
func (h *HTTPRequestHandler) userSignUpWithEmail(ctx *gin.Context) {
	var input signUpWithEmailInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	token, err := h.services.CreateUserByEmail(input.Email)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while creating user"))
		return
	}

	ctx.JSON(http.StatusCreated, dataResponse{Data: token})
}

func (h *HTTPRequestHandler) userRestorePassword(ctx *gin.Context) {

}

// @Summary Check user token and return user info
// @Tags auth
// @Description check token and return info if correct
// @Accept  json
// @Produce  json
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/check [get]
func (h *HTTPRequestHandler) userCheckToken(ctx *gin.Context) {
	h.isUserAuthenticated(ctx)

	_, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}
