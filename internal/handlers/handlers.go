package handlers

import (
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	httpHandler := http_1_1.NewRequestHandler(h.services)
	root := router.Group("/")
	{
		root.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello from server")
		})
		httpHandler.Init(root)
	}
}
