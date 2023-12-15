package order

import (
	"github.com/Le0nar/orders/internal/domain"
	orderCache "github.com/Le0nar/orders/internal/repository/cache/order"
	"github.com/Le0nar/orders/internal/repository/postgresql/order"
)

type OrderService struct{
	orderPsqlRepository order.OrderPsqlRepository
	orderCacheRepository orderCache.OrderCacheRepository
}

func NewOrderSerivce(or order.OrderPsqlRepository) *OrderService {
	return &OrderService{orderPsqlRepository: or}
}

func (os *OrderService) GetOrderById (id string) (string, error) {
	return os.orderPsqlRepository.GetOrderById(id)
}
func (os *OrderService) CreateOrder (order domain.Order)  error {
	// cache
	os.orderCacheRepository.CreateOrder(order)

	// db
	return os.orderPsqlRepository.CreateOrder(order)
}

