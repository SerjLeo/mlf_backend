package http_1_1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	GetCurrenciesRoute = ""
)

func (h *HTTPRequestHandler) initCurrencyRoutes(api *gin.RouterGroup) {
	acc := api.Group("/currency")
	{
		acc.GET(GetCurrenciesRoute, h.getCurrenciesList)
	}
}

func (h *HTTPRequestHandler) getCurrenciesList(ctx *gin.Context) {
	currencies, err := h.services.GetCurrenciesList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.Wrap(err, "error while getting currencies list"))
	}

	ctx.JSON(http.StatusOK, dataResponse{Data: currencies})
}
