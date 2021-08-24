package http_1_1

import (
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	services *services.Service
}

func NewRequestHandler(services *services.Service) *RequestHandler {
	return &RequestHandler{services: services}
}

func (r *RequestHandler) Init(root *gin.RouterGroup) {
	api := root.Group("/api")
	{
		r.initUserRoutes(api)
	}
}