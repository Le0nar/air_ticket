package repository

import (
	"fmt"

	orderCache "github.com/Le0nar/orders/internal/repository/cache/order"
	"github.com/Le0nar/orders/internal/repository/postgresql/order"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	OrderPsqlRepository *order.OrderPsqlRepository
	OrderCacheRepository *orderCache.OrderCacheRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderPsqlRepository: order.NewOrderPsqlRepository(db),
		OrderCacheRepository: orderCache.NewOrderCacheRepository(),
	}
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// TODO: mb move function to another file within the "repository" module
func NewPostgresDB(cfg DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
