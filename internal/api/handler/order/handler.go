package handler

import (
	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/google/uuid"
)

type OrderSerivce interface {
	GetOrderById(id uuid.UUID) (domain.Order, error)
}

type OrderHandler struct {
	OrderSerivce OrderSerivce
}

func NewOrderHandler(os OrderSerivce) *OrderHandler {
	return &OrderHandler{
		OrderSerivce: os,
	}
}