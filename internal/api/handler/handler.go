package handler

import (
	"github.com/Le0nar/orders/internal/api/handler/order"
	"github.com/Le0nar/orders/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	OrderHandler *order.OrderHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		OrderHandler: order.NewOrderHandler(*service.OrderService),
	}
}

func (h *Handler)  InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		order := api.Group("/order")
		{
			order.GET("/:id", h.OrderHandler.GetOrderById)
			order.POST("/", h.OrderHandler.CreateOrder)
		}
	}

	return router
}
