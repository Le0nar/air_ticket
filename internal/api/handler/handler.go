package handler

import (
	"github.com/Le0nar/orders/internal/api/handler/order"
	"github.com/Le0nar/orders/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	OrderHandler *order.OrderHandler
}

// type Service interface {
// 	order.OrderSerivce
// } 

func NewHandler(service *service.Service) *Handler {
	return &Handler{
	}
}

func (h *Handler)  InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		order := api.Group("/order")
		{
			order.GET("/:id", h.OrderHandler.GetOrderById)
		}
	}

	return router
}
