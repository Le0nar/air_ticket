package order

import (
	"fmt"

	domain "github.com/Le0nar/orders/internal/domain"
	"github.com/jmoiron/sqlx"
)

type OrderPsqlRepository struct {
	db *sqlx.DB
}

func NewOrderPsqlRepository(db *sqlx.DB) *OrderPsqlRepository {
	return &OrderPsqlRepository{db: db}
}

const ordersTable = "orders"

func (opr *OrderPsqlRepository) GetOrderById (id string) (string, error) {
	var order domain.Order

	query := fmt.Sprintf("SELECT * FROM %s where id = '%s';", ordersTable, id)
	fmt.Printf("query: %v\n", query)
	err := opr.db.Get(&order, query)

	return order.Content, err
}

func (opr *OrderPsqlRepository) CreateOrder(order domain.Order)  error {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (id, content) values ($1, $2) RETURNING id", ordersTable)

	row := opr.db.QueryRow(query, order.Id, order.Content)
	if err := row.Scan(&id); err != nil {
		return  err
	}

	return  nil
}

