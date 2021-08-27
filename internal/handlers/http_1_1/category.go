package http_1_1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// @Summary Get categories list
// @Tags category
// @Description returns user categories list with pagination
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param input body metaParams false "pagination params and filters"
// @Success 200 {object} dataWithPaginationResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category [get]
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
		data: categories,
	})
}

// @Summary Get category by id
// @Tags category
// @Description returns user's category object by id
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param categoryId path integer false "target category id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category/{categoryId} [get]
func (h *RequestHandler) getCategoryById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized,err.Error())
		return
	}

	categoryId := c.Param("id")
	if categoryId == "" {
		newErrorResponse(c,http.StatusBadRequest,"category id is not provided")
		return
	}

	catId, err := strconv.Atoi(categoryId)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.services.Category.GetUserCategoryById(userId, catId)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprint("Error getting category:",err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		data: category,
	})
}

func (h *RequestHandler) createCategory(c *gin.Context) {

}

func (h *RequestHandler) updateCategory(c *gin.Context) {

}

func (h *RequestHandler) deleteCategory(c *gin.Context) {

}