package http_1_1

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type HTTPRequestHandler struct {
	services handlers.Service
}

func NewRequestHandler(services handlers.Service) *HTTPRequestHandler {
	return &HTTPRequestHandler{services: services}
}

func (h *HTTPRequestHandler) InitRoutes() *gin.Engine {
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

func (h *HTTPRequestHandler) initApi(router *gin.Engine) {
	root := router.Group("/")
	{
		root.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello from server")
		})
		root.GET("/send-email", func(context *gin.Context) {
			err := h.services.SendTestEmail()
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

		api := root.Group("/api")
		{
			h.initAuthRoutes(api)
			h.initCategoriesRoutes(api)
			h.initTransactionsRoutes(api)
			h.initUserRoutes(api)
		}
	}
}
