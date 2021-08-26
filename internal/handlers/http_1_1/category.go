package http_1_1

import "github.com/gin-gonic/gin"

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

}

func (h *RequestHandler) getCategoryById(c *gin.Context) {

}

func (h *RequestHandler) createCategory(c *gin.Context) {

}

func (h *RequestHandler) updateCategory(c *gin.Context) {

}

func (h *RequestHandler) deleteCategory(c *gin.Context) {

}