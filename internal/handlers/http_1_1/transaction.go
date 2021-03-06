package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
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
// @Param input body metaParams false "pagination params and filters"
// @Success 200 {object} dataWithPaginationResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transactions [get]
func (h *HTTPRequestHandler) getTransactionsList(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	transactions, err := h.services.GetTransactions(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("Error getting transactions list:", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
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
func (h *HTTPRequestHandler) getTransactionById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId := c.Param("id")
	if transactionId == "" {
		newErrorResponse(c, http.StatusBadRequest, "transaction id is not provided")
		return
	}

	transId, err := strconv.Atoi(transactionId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction, err := h.services.GetTransactionById(userId, transId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("Error getting transaction:", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

// @Summary Create new transaction
// @Tags transaction
// @Description creates new transaction and returns it
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param input body models.CreateTransactionInput true "created transaction fields"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transactions [post]
func (h *HTTPRequestHandler) createTransaction(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input models.CreateTransactionInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid data for create").Error())
		return
	}

	fmt.Printf("%+v\n", input)

	transaction, err := h.services.CreateTransaction(userId, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("Error while creating transaction: ", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

func (h *HTTPRequestHandler) updateTransaction(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id").Error())
		return
	}

	var input models.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transaction, err := h.services.UpdateTransaction(userId, transactionId, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("Error while updating transaction: ", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: transaction,
	})
}

func (h *HTTPRequestHandler) deleteTransaction(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id").Error())
		return
	}

	if err := h.services.DeleteTransaction(userId, transactionId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("error while deleting transaction: ", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: transactionId,
	})
}

func (h *HTTPRequestHandler) attachCategories(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id").Error())
		return
	}

	categoryId, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid category id").Error())
		return
	}

	err = h.services.AttachCategory(userId, transactionId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("error while attach category:", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}

func (h *HTTPRequestHandler) detachCategories(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid transaction id").Error())
		return
	}

	categoryId, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, errors.Wrap(err, "invalid category id").Error())
		return
	}

	err = h.services.DetachCategory(userId, transactionId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprint("error while detach category:", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}
