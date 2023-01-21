package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	GetAccountsRoute    = "/"
	GetAccountByIdRoute = "/:id"
	CreateAccountRoute  = "/"
	UpdateAccountRoute  = "/:id"
	DeleteAccountRoute  = "/:id"
)

func (h *HTTPRequestHandler) initAccountRoutes(api *gin.RouterGroup) {
	auth := api.Group("/account", h.isUserAuthenticated)
	{
		auth.GET(GetAccountsRoute, h.withPagination, h.GetUserAccounts)
		auth.GET(GetAccountByIdRoute, h.GetUserAccountById)
		auth.POST(CreateAccountRoute, h.CreateUserAccount)
		auth.PUT(UpdateAccountRoute, h.UpdateUserAccount)
		auth.DELETE(DeleteAccountRoute, h.DeleteUserAccount)
	}
}

func (h *HTTPRequestHandler) GetUserAccounts(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	pagination, err := h.getPagination(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	accounts, err := h.services.GetAccounts(pagination, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error getting accounts:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: accounts,
	})
}

func (h *HTTPRequestHandler) GetUserAccountById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "account id is not provided")
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.services.GetAccountById(accId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error getting account:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) CreateUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var input models.CreateAccountInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid data for create").Error())
		return
	}

	account, err := h.services.CreateAccount(&input, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error getting account:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) UpdateUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "account id is not provided")
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateAccountInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "invalid data for update").Error())
		return
	}

	account, err := h.services.UpdateAccount(accId, userId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error updating account:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) DeleteUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "account id is not provided")
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.SoftDeleteAccount(accId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprint("Error getting account:", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}
