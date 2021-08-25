package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RequestHandler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-in", r.userSignIn)
		user.POST("/sign-up", r.userSignUp)
		user.POST("/refresh", r.userRefreshAccess)
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RequestHandler) userSignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}
}

func (r *RequestHandler) userSignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %s", err.Error()))
		return
	}

	token, err := r.services.User.Create(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error while creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dataResponse{Data: token})

}

func (r *RequestHandler) userRefreshAccess(c *gin.Context) {

}
