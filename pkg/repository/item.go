package repository

import (
	"L0/pkg/repository/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(orderId string, items []*models.Item) ([]*models.Item, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to beginx in items creating: %v", err)
	}

	createdItems := make([]*models.Item, 0, len(items))
	for _, item := range items {
		var i models.Item
		createItemQuery := fmt.Sprintf("INSERT INTO %s (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *", itemsTable)

		row := tx.QueryRowx(createItemQuery, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
		if err = row.StructScan(&i); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to scan struct in items creating: %v", err)
		}

		createOrdersItemsQuery := fmt.Sprintf("INSERT INTO %s (order_uid, item_id) VALUES ($1, $2)", ordersItemsTable)
		_, err = tx.Exec(createOrdersItemsQuery, orderId, &i.ItemId)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to exec in items creating: %v", err)
		}
		createdItems = append(createdItems, &i)
	}

	return createdItems, tx.Commit()
}

func (r *ItemPostgres) Get(orderId string) ([]*models.Item, error) {
	query := fmt.Sprintf("SELECT i.* FROM %s i INNER JOIN %s oi on i.item_id = oi.item_id WHERE oi.order_uid = $1",
		itemsTable, ordersItemsTable)
	rows, err := r.db.Queryx(query, orderId)
	if err != nil {
		return nil, fmt.Errorf("cant create query to get items: %w", err)
	}

	items := make([]*models.Item, 0)

	for rows.Next() {
		var item models.Item
		err = rows.StructScan(&item)
		if err != nil {
			return nil, fmt.Errorf("failed to scan items in items get: %w", err)
		}
		items = append(items, &item)
	}

	return items, nil
}

//func (r *ItemPostgres) GetById(orderId string, itemId int) (*models.Item, error) {
//	var item models.Item
//
//	query := fmt.Sprintf("SELECT i.* FROM %s i INNER JOIN %s oi on i.item_id = oi.item_id WHERE oi.order_uid = $1 AND oi.item_id = $2",
//		itemsTable, ordersItemsTable)
//	row := r.db.QueryRowx(query, orderId, itemId)
//	if err := row.StructScan(&item); err != nil {
//		return nil, fmt.Errorf("failed to scan item in items get by id: %v", err)
//	}
//
//	return &item, nil
//}
