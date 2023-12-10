package order

import (
	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/Le0nar/orders/internal/repository/order"
	"github.com/google/uuid"
)

// type OrderRepository interface {
// 	GetOrderById(id uuid.UUID) (domain.Order, error)
// }

type OrderService struct{
	orderRepository order.OrderRepository
}

func NewOrderSerivce(or order.OrderRepository) *OrderService {
	return &OrderService{orderRepository: or}
}

func (os *OrderService) GetOrderById (id uuid.UUID) (domain.Order, error) {
	return os.orderRepository.GetOrderById(id)
}