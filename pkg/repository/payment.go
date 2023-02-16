package repository

import (
	"L0/pkg/repository/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) Create(orderId string, payment *models.Payment) (*models.Payment, error) {
	var pay models.Payment

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to beginx in payment creating: %v", err)
	}

	createPaymentQuery := fmt.Sprintf("INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *", paymentsTable)
	row := tx.QueryRowx(createPaymentQuery, payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err = row.StructScan(&pay); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to scan struct in payment's creating: %v", err)
	}

	createOrdersPaymentQuery := fmt.Sprintf("INSERT INTO %s (order_uid, payment_id) VALUES ($1, $2)", ordersPaymentTable)
	_, err = tx.Exec(createOrdersPaymentQuery, orderId, pay.PaymentId)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to exec in payment's creating: %v", err)
	}

	return &pay, tx.Commit()
}

func (r *PaymentPostgres) Get(orderId string) (*models.Payment, error) {
	var payment models.Payment

	query := fmt.Sprintf("SELECT p.* FROM %s p INNER JOIN %s op on p.payment_id = op.payment_id WHERE op.order_uid = $1",
		paymentsTable, ordersPaymentTable)
	row := r.db.QueryRowx(query, orderId)
	if err := row.StructScan(&payment); err != nil {
		return nil, fmt.Errorf("failed to scan payment in payment's get: %v", err)
	}

	return &payment, nil
}
