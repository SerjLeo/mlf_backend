package handlers

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

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
		cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     false,
			ValidateHeaders: false,
		}),
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
		root.GET("/send-email", func(context *gin.Context) {
			err := h.services.User.SendTestEmail()
			if err != nil {
				fmt.Println(err.Error())
				context.JSON(http.StatusInternalServerError, map[string]interface{}{
					"err": err.Error(),
				})
			}
			return
			context.String(http.StatusOK, "Email sent")
		})
		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		httpHandler.Init(root)
	}
}
