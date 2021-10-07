package http_1_1

import "github.com/gin-gonic/gin"

func (h *RequestHandler) initTransactionsRoutes(api *gin.RouterGroup) {
	transaction := api.Group("/transaction", h.isUserAuthenticated)
	{
		transaction.GET("/", h.getTransactionsList)
		transaction.GET("/:id", h.getTransactionById)
		transaction.POST("/", h.createTransaction)
		transaction.PUT("/:id", h.updateTransaction)
		transaction.DELETE("/:id", h.deleteTransaction)
		transaction.POST("/:id/category-attach", h.attachCategories)
		transaction.POST("/:id/category-detach", h.detachCategories)
	}
}

func (h *RequestHandler) getTransactionsList(c *gin.Context) {}

func (h *RequestHandler) getTransactionById(c *gin.Context) {}

func (h *RequestHandler) createTransaction(c *gin.Context) {}

func (h *RequestHandler) updateTransaction(c *gin.Context) {}

func (h *RequestHandler) deleteTransaction(c *gin.Context) {}

func (h *RequestHandler) attachCategories(c *gin.Context) {}

func (h *RequestHandler) detachCategories(c *gin.Context) {}