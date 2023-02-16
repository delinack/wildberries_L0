package repository

import (
	"L0/pkg/repository/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DeliveryPostgres struct {
	db *sqlx.DB
}

func NewDeliveryPostgres(db *sqlx.DB) *DeliveryPostgres {
	return &DeliveryPostgres{db: db}
}

func (r *DeliveryPostgres) Create(orderId string, delivery *models.Delivery) (*models.Delivery, error) {
	var del models.Delivery

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to beginx in delivery creating: %v", err)
	}

	createDeliveryQuery := fmt.Sprintf("INSERT INTO %s (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", deliveriesTable)
	row := tx.QueryRowx(createDeliveryQuery, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err = row.StructScan(&del); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to scan struct in delivery creating: %v", err)
	}

	createOrdersDeliveryQuery := fmt.Sprintf("INSERT INTO %s (order_uid, delivery_id) VALUES ($1, $2)", ordersDeliveryTable)
	_, err = tx.Exec(createOrdersDeliveryQuery, orderId, del.DeliveryId)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to exec in delivery creating: %v", err)
	}

	return &del, tx.Commit()
}

func (r *DeliveryPostgres) Get(orderId string) (*models.Delivery, error) {
	var delivery models.Delivery

	query := fmt.Sprintf("SELECT d.* FROM %s d INNER JOIN %s od on d.delivery_id = od.delivery_id WHERE od.order_uid = $1", deliveriesTable, ordersDeliveryTable)
	row := r.db.QueryRowx(query, orderId)
	if err := row.StructScan(&delivery); err != nil {
		return nil, fmt.Errorf("failed to scan delivery in delivery's get: %v", err)
	}

	return &delivery, nil
}
