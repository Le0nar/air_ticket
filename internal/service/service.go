package service

import (
	"github.com/Le0nar/orders/internal/repository"
	"github.com/Le0nar/orders/internal/service/order"
)

type Service struct {
	OrderService *order.OrderService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		OrderService: order.NewOrderSerivce(*repo.OrderRepository),
	}
}
