package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RequestHandler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.GET("/check", h.userCheckToken)
		auth.POST("/sign-in", h.userSignIn)
		auth.POST("/sign-up", h.userSignUp)
		auth.POST("/sign-up-with-email", h.userSignUpWithEmail)
		auth.POST("/restore-password", h.userRestorePassword)
	}
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary User sign-in with email and password
// @Tags auth
// @Description returns auth JWT
// @Accept  json
// @Produce  json
// @Param input body signInInput true "info for user's login"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *RequestHandler) userSignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}

	token, err := h.services.User.SignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: token})
}

type signUpInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary User sign-up with name, email and password
// @Tags auth
// @Description creates user and returns auth JWT
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "data for user creation"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *RequestHandler) userSignUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}

	token, err := h.services.User.Create(models.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error while creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dataResponse{Data: token})

}

type signUpWithEmailInput struct {
	Email string `json:"email" binding:"required"`
}

// @Summary User sign-up with email only
// @Tags auth
// @Description creates user with email, password generates automatically and returns auth JWT
// @Accept  json
// @Produce  json
// @Param input body signUpWithEmailInput true "email for user creation"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up-with-email [post]
func (h *RequestHandler) userSignUpWithEmail(c *gin.Context) {
	var input signUpWithEmailInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}

	token, err := h.services.User.CreateByEmail(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error while creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dataResponse{Data: token})
}

func (h *RequestHandler) userRestorePassword(c *gin.Context) {

}

func (h *RequestHandler) userCheckToken(c *gin.Context) {
	h.isUserAuthenticated(c)

	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while checking authorization")
		return
	}

	user, err := h.services.User.GetUserProfile(userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user profile doesn't exists")
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: userResponse{Name: user.Name, Email: user.Email}})
}
