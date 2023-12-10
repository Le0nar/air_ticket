package order

import (
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

func (or *OrderRepository) GetOrderById (id string) (string, error) {
	var order domain.Order

	query := fmt.Sprintf("SELECT * FROM %s where id = '%s';", ordersTable, id)
	fmt.Printf("query: %v\n", query)
	err := or.db.Get(&order, query)

	return order.Content, err
}
