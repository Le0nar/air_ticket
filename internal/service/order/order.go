package order

import "github.com/Le0nar/orders/internal/repository/postgresql/order"

type OrderService struct{
	orderPsqlRepository order.OrderPsqlRepository
}

func NewOrderSerivce(or order.OrderPsqlRepository) *OrderService {
	return &OrderService{orderPsqlRepository: or}
}

func (os *OrderService) GetOrderById (id string) (string, error) {
	return os.orderPsqlRepository.GetOrderById(id)
}