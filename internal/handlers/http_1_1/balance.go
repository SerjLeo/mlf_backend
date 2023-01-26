package http_1_1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	GetTotalBalancesRoute = "/total"
)

func (h *HTTPRequestHandler) initBalanceRoutes(api *gin.RouterGroup) {
	acc := api.Group("/balance", h.isUserAuthenticated)
	{
		acc.GET(GetTotalBalancesRoute, h.getTotalBalances)
	}
}

func (h *HTTPRequestHandler) getTotalBalances(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	balances, err := h.services.GetUserBalancesAmount(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, errors.Wrap(err, "error while getting balances"))
		return
	}

	ctx.JSON(http.StatusOK, dataResponse{
		Data: balances,
	})
}
