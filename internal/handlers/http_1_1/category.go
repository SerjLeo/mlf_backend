package http_1_1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RequestHandler) initCategoriesRoutes(api *gin.RouterGroup) {
	category := api.Group("/category", h.isUserAuthenticated)
	{
		category.GET("/", h.getCategoriesList)
		category.GET("/:id", h.getCategoryById)
		category.POST("/", h.createCategory)
		category.PUT("/:id", h.updateCategory)
		category.DELETE("/:id", h.deleteCategory)
	}
}

func (h *RequestHandler) getCategoriesList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized,err.Error())
		return
	}

	categories, err := h.services.Category.GetUserCategories(userId)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprint("Error getting categories list:",err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: categories,
	})
}

func (h *RequestHandler) getCategoryById(c *gin.Context) {

}

func (h *RequestHandler) createCategory(c *gin.Context) {

}

func (h *RequestHandler) updateCategory(c *gin.Context) {

}

func (h *RequestHandler) deleteCategory(c *gin.Context) {

}