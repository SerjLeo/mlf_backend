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
	GetAccountsRoute    = ""
	GetAccountByIdRoute = "/:id"
	CreateAccountRoute  = ""
	UpdateAccountRoute  = "/:id"
	DeleteAccountRoute  = "/:id"
)

func (h *HTTPRequestHandler) initAccountRoutes(api *gin.RouterGroup) {
	acc := api.Group("/account", h.isUserAuthenticated)
	{
		acc.GET(GetAccountsRoute, h.withPagination, h.getUserAccounts)
		acc.GET(GetAccountByIdRoute, h.getUserAccountById)
		acc.POST(CreateAccountRoute, h.createUserAccount)
		acc.PUT(UpdateAccountRoute, h.updateUserAccount)
		acc.DELETE(DeleteAccountRoute, h.deleteUserAccount)
	}
}

func (h *HTTPRequestHandler) getUserAccounts(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	pagination, err := h.getPagination(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	accounts, err := h.services.GetAccounts(pagination, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error getting accounts"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: accounts,
	})
}

func (h *HTTPRequestHandler) getUserAccountById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountIDIsNotProvided)
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountInvalidID)
		return
	}

	account, err := h.services.GetAccountById(accId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) createUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	var input models.CreateAccountInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	account, err := h.services.CreateAccount(&input, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while creating account"))
		return
	}

	ctx.JSON(http.StatusCreated, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) updateUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountIDIsNotProvided)
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountInvalidID)
		return
	}

	var input models.UpdateAccountInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.BadInput)
		return
	}

	account, err := h.services.UpdateAccount(accId, userId, &input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while updating account"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: account,
	})
}

func (h *HTTPRequestHandler) deleteUserAccount(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	accountId := ctx.Param("id")
	if accountId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountIDIsNotProvided)
		return
	}

	accId, err := strconv.Atoi(accountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, custom_errors.AccountInvalidID)
		return
	}

	err = h.services.SoftDeleteAccount(accId, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: true,
	})
}
