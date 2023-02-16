package handler

import (
	"L0/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	*service.HTTPService
}

func NewHandler(services *service.HTTPService) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	order := router.Group("/orders")
	{
		order.GET("/:id", h.getOrderById)
		order.GET("/api/:id", h.getJSONOrderById)
	}

	return router
}
