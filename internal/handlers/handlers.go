package handlers

import (
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/SerjLeo/mlf_backend/docs"
)

type Handler struct {
	services     *services.Service
	tokenManager auth.TokenManager
}

func NewHandler(services *services.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{services: services, tokenManager: tokenManager}
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
		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		httpHandler.Init(root)
	}
}
