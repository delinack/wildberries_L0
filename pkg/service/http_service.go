package service

import (
	"L0/pkg/repository"
)

type HTTPService struct {
	OrderService Order
}

func NewHTTPService(repo *repository.Repository) *HTTPService {
	return &HTTPService{
		OrderService: NewOrderService(repo),
	}
}
