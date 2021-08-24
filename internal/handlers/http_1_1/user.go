package http_1_1

import (
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

func (r *RequestHandler) userSignIn(c *gin.Context) {
	c.String(http.StatusOK, "Signing in...")
}

func (r *RequestHandler) userSignUp(c *gin.Context) {
}

func (r *RequestHandler) userRefreshAccess(c *gin.Context) {

}
