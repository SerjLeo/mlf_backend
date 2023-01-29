package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/models/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (h *HTTPRequestHandler) initTransactionsRoutes(api *gin.RouterGroup) {
	transaction := api.Group("/transaction", h.isUserAuthenticated)
	{
		transaction.GET("", h.getTransactionsList)
		transaction.GET("/:id", h.getTransactionById)
		transaction.POST("", h.createTransaction)
		transaction.PUT("/:id", h.updateTransaction)
		transaction.DELETE("/:id", h.deleteTransaction)
		transaction.POST("/:id/category-attach/:category_id", h.attachCategories)
		transaction.POST("/:id/category-detach/:category_id", h.detachCategories)
	}
}

// @Summary Get transactions list
// @Tags transaction
// @Description returns transactions list with pagination
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body metaParams false "pagination params and filters"
// @Success 200 {object} dataWithPaginationResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transactions [get]
func (h *HTTPRequestHandler) getTransactionsList(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactions, err := h.services.GetTransactions(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting transactions list"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: transactions,
	})
}

// @Summary Get transaction by id
// @Tags transaction
// @Description returns user's transaction object by id
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param categoryId path integer false "target transaction id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transactions/{transactionId} [get]
func (h *HTTPRequestHandler) getTransactionById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactionId := ctx.Param("id")
	if transactionId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.TransactionIdNotProvided)
		return
	}

	transId, err := strconv.Atoi(transactionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	transaction, err := h.services.GetTransactionById(userId, transId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting transaction"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

// @Summary Create new transaction
// @Tags transaction
// @Description creates new transaction and returns it
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body models.CreateTransactionInput true "created transaction fields"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transactions [post]
func (h *HTTPRequestHandler) createTransaction(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	var input models.CreateTransactionInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid data for create"))
		return
	}

	fmt.Printf("%+v\n", input)

	transaction, err := h.services.CreateTransaction(userId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while creating transaction"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

func (h *HTTPRequestHandler) updateTransaction(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id"))
		return
	}

	var input models.Transaction

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, custom_errors.BadInput)
		return
	}

	transaction, err := h.services.UpdateTransaction(userId, transactionId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while updating transaction"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

func (h *HTTPRequestHandler) deleteTransaction(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id"))
		return
	}

	if err := h.services.DeleteTransaction(userId, transactionId); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while deleting category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: transactionId,
	})
}

func (h *HTTPRequestHandler) attachCategories(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id"))
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid category id"))
		return
	}

	err = h.services.AttachCategory(userId, transactionId, categoryId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while attach category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}

func (h *HTTPRequestHandler) detachCategories(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id"))
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid category id"))
		return
	}

	err = h.services.DetachCategory(userId, transactionId, categoryId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while detach category"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}
