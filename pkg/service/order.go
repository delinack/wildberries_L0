package service

import (
	"L0/pkg/domain"
	"L0/pkg/repository"
	"L0/pkg/repository/models"
	"fmt"
)

type Order interface {
	CreateOrder(order *domain.Order) (*models.Order, error)
	GetByIdFromCache(orderId string) (*domain.Order, error)
	PullOrders() error
}

type OrderService struct {
	repo *repository.Repository

	cacheMap map[string]*domain.Order
}

func NewOrderService(repo *repository.Repository) Order {
	return &OrderService{repo: repo}
}

func (s *OrderService) PullOrders() error {
	m := make(map[string]*domain.Order, 0)

	temp, err := s.repo.Order.Get()
	if err != nil {
		return fmt.Errorf("repo.Order.Get error in PullOrders: %w", err)
	}

	for _, order := range temp {
		payment, err := s.repo.Payment.Get(order.OrderUid)
		if err != nil {
			return fmt.Errorf("paymentRepository.Get error PullOrders: %w", err)
		}

		delivery, err := s.repo.Delivery.Get(order.OrderUid)
		if err != nil {
			return fmt.Errorf("deliveryRepository.Get error PullOrders: %w", err)
		}

		items, err := s.repo.Item.Get(order.OrderUid)
		if err != nil {
			return fmt.Errorf("itemRepository.Get error PullOrders: %w", err)
		}

		m[order.OrderUid] = ConvertOrderToDomain(order, delivery, payment, items)
	}
	s.cacheMap = m

	return nil
}

func (s *OrderService) CreateOrder(order *domain.Order) (*models.Order, error) {
	m := make(map[string]*domain.Order, 0)

	createdOrder, err := s.repo.Order.CreateOrder(ConvertOrderToModel(order))
	if err != nil {
		return nil, fmt.Errorf("orderRepository.CreateOrder error: %w", err)
	}

	payment, err := s.repo.Payment.Create(createdOrder.OrderUid, ConvertPaymentToModel(order.Payment))
	if err != nil {
		return nil, fmt.Errorf("paymentRepository.Create error: %w", err)
	}

	delivery, err := s.repo.Delivery.Create(createdOrder.OrderUid, ConvertDeliveryToModel(order.Delivery))
	if err != nil {
		return nil, fmt.Errorf("deliveryRepository.Create error: %w", err)
	}

	items, err := s.repo.Item.Create(createdOrder.OrderUid, ConvertItemsToModel(order.Items))
	if err != nil {
		return nil, fmt.Errorf("itemRepository.Create error: %w", err)
	}

	m[createdOrder.OrderUid] = ConvertOrderToDomain(createdOrder, delivery, payment, items)
	s.cacheMap = m

	return createdOrder, nil
}

func (s *OrderService) GetByIdFromCache(orderId string) (*domain.Order, error) { //////
	if _, ok := s.cacheMap[orderId]; !ok {
		return nil, fmt.Errorf("%s id does not exist", orderId)
	}

	return s.cacheMap[orderId], nil
}
