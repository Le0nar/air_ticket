package order

import (
	"encoding/json"
	"fmt"

	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

const ordersTable = "orders"

func (or *OrderRepository) GetOrderById (id string) (domain.Order, error) {
	var orderDTO domain.OrderDTO
	query := fmt.Sprintf("SELECT * FROM %s where id = '%s';", ordersTable, id)
	fmt.Printf("query: %v\n", query)
	err := or.db.Get(&orderDTO, query)

	var order domain.Order
	json.Unmarshal([]byte(orderDTO.Content), &order)

	return order, err
}
