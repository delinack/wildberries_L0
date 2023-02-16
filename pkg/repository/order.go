package repository

import (
	"L0/pkg/repository/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(order *models.Order) (*models.Order, error) {
	var or models.Order

	createOrderQuery := fmt.Sprintf("INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *", ordersTable)
	row := r.db.QueryRowx(createOrderQuery, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err := row.StructScan(&or); err != nil {
		return nil, fmt.Errorf("failed to scan order in create order: %v", err)
	}

	return &or, nil
}

func (r *OrderPostgres) GetById(orderId string) (*models.Order, error) {
	var order models.Order

	query := fmt.Sprintf("SELECT * FROM %s o WHERE o.order_uid = $1", ordersTable)
	row := r.db.QueryRowx(query, orderId)
	if err := row.StructScan(&order); err != nil {
		return nil, fmt.Errorf("failed to scan order in get order by id: %v", err)
	}

	return &order, nil
}

func (r *OrderPostgres) Get() ([]*models.Order, error) {
	query := fmt.Sprintf("SELECT * FROM %s", ordersTable)
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("cant create query to get items: %w", err)
	}

	orders := make([]*models.Order, 0)

	for rows.Next() {
		var order models.Order
		err = rows.StructScan(&order)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order in orders get: %w", err)
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
