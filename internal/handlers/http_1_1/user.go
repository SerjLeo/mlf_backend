package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RequestHandler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-in", h.userSignIn)
		user.POST("/sign-up", h.userSignUp)
		user.POST("/sign-up-with-email", h.userSignUpWithEmail)
		user.POST("/refresh", h.userRefreshAccess)
	}
}

type signInInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

func (h *RequestHandler) userSignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}

	token, err := h.services.User.Create(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error while creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dataResponse{Data: token})

}

type signUpWithEmailInput struct {
	Email string `json:"email" binding:"required"`
}

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

func (h *RequestHandler) userRefreshAccess(c *gin.Context) {

}
