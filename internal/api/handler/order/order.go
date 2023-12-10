package order

import (
	"net/http"

	"github.com/Le0nar/orders/internal/service/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type OrderSerivce interface {
// 	GetOrderById(id uuid.UUID) (domain.Order, error)
// }

type OrderHandler struct {
	OrderSerivce order.OrderService
}

func NewOrderHandler(os order.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderSerivce: os,
	}
}

func (or *OrderHandler) GetOrderById (c *gin.Context) {
	id := c.Param("id")
	
	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	news, err := or.OrderSerivce.GetOrderById(parsedId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, news)
}
