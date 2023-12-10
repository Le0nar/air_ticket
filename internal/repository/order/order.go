package order

import (
	"fmt"

	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

const ordersTable = "orders"

func (or *OrderRepository) GetOrderById (id uuid.UUID) (domain.Order, error) {
	var order domain.Order
	query := fmt.Sprintf("SELECT * FROM %s where id = %d", ordersTable, id)
	err := or.db.Get(&order, query)

	return order, err
}
