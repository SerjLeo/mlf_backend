package http_1_1

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

const (
	GetCategoriesRoute   = ""
	GetCategoryByIdRoute = "/:id"
	UpdateCategoryRoute  = "/:id"
	CreateCategoryRoute  = ""
	DeleteCategoryRoute  = "/:id"
)

func (h *HTTPRequestHandler) initCategoriesRoutes(api *gin.RouterGroup) {
	category := api.Group("/category", h.isUserAuthenticated)
	{
		category.GET(GetCategoriesRoute, h.withPagination, h.getCategoriesList)
		category.GET(GetCategoryByIdRoute, h.getCategoryById)
		category.POST(CreateCategoryRoute, h.createCategory)
		category.PUT(UpdateCategoryRoute, h.updateCategory)
		category.DELETE(DeleteCategoryRoute, h.deleteCategory)
	}
}

// @Summary Get categories list
// @Tags category
// @Description returns user categories list with pagination
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body metaParams false "pagination params and filters"
// @Success 200 {object} dataWithPaginationResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category [get]
func (h *HTTPRequestHandler) getCategoriesList(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	pagination, err := h.getPagination(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	categories, err := h.services.GetUserCategories(userId, pagination)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting categories"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: categories,
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
func (h *HTTPRequestHandler) getCategoryById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	categoryId := ctx.Param("id")
	if categoryId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.CategoryIdNotProvided)
		return
	}

	catId, err := strconv.Atoi(categoryId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.CategoryInvalidID)
		return
	}

	category, err := h.services.GetUserCategoryById(userId, catId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: category,
	})
}

// @Summary Create new category
// @Tags category
// @Description creates new category and returns it
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body models.CreateCategoryInput true "created category fields"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category [post]
func (h *HTTPRequestHandler) createCategory(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	var input models.CreateCategoryInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, custom_errors.BadInput)
		return
	}

	category, err := h.services.CreateCategory(userId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while creating category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: category,
	})
}

// @Summary Update existing category
// @Tags category
// @Description updates existing category and returns updated instance
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body models.Category true "updated category fields"
// @Param categoryId path integer false "target category id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category/{categoryId} [put]
func (h *HTTPRequestHandler) updateCategory(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.CategoryInvalidID)
		return
	}

	var input models.UpdateCategoryInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	category, err := h.services.UpdateCategory(userId, categoryId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while updating category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: category,
	})
}

// @Summary Delete existing category
// @Tags category
// @Description delete existing category and returns id of deleted category
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param categoryId path integer false "target category id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /category/{categoryId} [delete]
func (h *HTTPRequestHandler) deleteCategory(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.CategoryInvalidID)
		return
	}

	if err := h.services.DeleteCategory(userId, categoryId); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while deleting category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: categoryId,
	})
}
