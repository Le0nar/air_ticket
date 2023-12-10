package order

import (
	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrderById(id uuid.UUID) (domain.Order, error)
}

type OrderService struct{
	orderRepository OrderRepository
}

func NewOrderSerivce(or OrderRepository) *OrderService {
	return &OrderService{orderRepository: or}
}
