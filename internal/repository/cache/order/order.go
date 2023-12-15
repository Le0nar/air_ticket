package order

import (
	"fmt"

	"github.com/Le0nar/orders/internal/domain"
)

type OrderCacheRepository struct {
	orderList []domain.Order
}

func NewOrderCacheRepository() *OrderCacheRepository {
	return &OrderCacheRepository{}
}

func (ocr *OrderCacheRepository) CreateOrder(order domain.Order) {
	newArray := append(ocr.orderList, order)

	

	fmt.Printf("newArray: %v\n", newArray)
}


// TODO: for recover data from postgresql
// func (ocr *OrderCacheRepository) CreateOrderList(orderList []domain.Order) {
// 	for _, order := range orderList {
// 		ocr.orderCache[order.Id] = order.Content
// 	}
// }
