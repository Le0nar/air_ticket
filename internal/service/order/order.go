package order

import (
	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/Le0nar/orders/internal/repository/order"
)

type OrderService struct{
	orderRepository order.OrderRepository
}

func NewOrderSerivce(or order.OrderRepository) *OrderService {
	return &OrderService{orderRepository: or}
}

func (os *OrderService) GetOrderById (id string) (domain.Order, error) {
	return os.orderRepository.GetOrderById(id)
}