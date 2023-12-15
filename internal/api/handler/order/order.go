package order

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Le0nar/orders/internal/api/handler/errmes"
	"github.com/Le0nar/orders/internal/domain"
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


type orderIdWrapper struct {
	OrderUID    string `json:"order_uid"`
}

func (oh *OrderHandler) CreateOrder(c *gin.Context) {
	bodyAsByteArray, err := io.ReadAll(c.Request.Body) 
	if err != nil {
		errmes.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var orderData map[string]interface{}

    if err := json.Unmarshal(bodyAsByteArray, &orderData); err != nil {
        errmes.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
    }

	id, ok := orderData["order_uid"]
	if !ok {
		errmes.NewErrorResponse(c, http.StatusBadRequest, "order_uid is absent")
		return
	}

	stringedId, ok := id.(string)
	if !ok {
		errmes.NewErrorResponse(c, http.StatusBadRequest, "order_uid is not string")
		return
	}

	order := domain.Order{Content: string(bodyAsByteArray), Id: stringedId}

	fmt.Printf("order: %v\n", order)


	err = oh.OrderSerivce.CreateOrder(order)
	if err != nil {
		errmes.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}
