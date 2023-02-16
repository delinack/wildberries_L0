package repository

import (
	"L0/pkg/repository/models"

	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetById(orderId string) (*models.Order, error)
	Get() ([]*models.Order, error)
}

type Payment interface {
	Create(orderId string, payment *models.Payment) (*models.Payment, error)
	Get(orderId string) (*models.Payment, error)
}

type Delivery interface {
	Create(orderId string, delivery *models.Delivery) (*models.Delivery, error)
	Get(orderId string) (*models.Delivery, error)
}

type Item interface {
	Create(orderId string, item []*models.Item) ([]*models.Item, error)
	Get(orderId string) ([]*models.Item, error)
}

type Repository struct {
	Order    Order
	Payment  Payment
	Delivery Delivery
	Item     Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order:    NewOrderPostgres(db),
		Payment:  NewPaymentPostgres(db),
		Delivery: NewDeliveryPostgres(db),
		Item:     NewItemPostgres(db),
	}
}
