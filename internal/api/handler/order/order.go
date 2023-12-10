package order

import (
	"net/http"

	"github.com/Le0nar/orders/internal/api/handler/errmes"
	"github.com/Le0nar/orders/internal/service/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
	
	_, err := uuid.Parse(id)
	if err != nil {
		errmes.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := or.OrderSerivce.GetOrderById(id)
	if err != nil {
		errmes.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, order)
}
