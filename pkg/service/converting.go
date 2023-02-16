package service

import (
	"L0/pkg/domain"
	"L0/pkg/repository/models"
)

func ConvertItemsToDomain(items []*models.Item) []*domain.Item {
	res := make([]*domain.Item, len(items))
	for i, item := range items {
		res[i] = ConvertItemToDomain(item)
	}

	return res
}

func ConvertItemToDomain(item *models.Item) *domain.Item {
	return &domain.Item{
		ChrtID:      item.ChrtId,
		TrackNumber: item.TrackNumber,
		Price:       item.Price,
		Rid:         item.Rid,
		Name:        item.Name,
		Sale:        item.Sale,
		Size:        item.Size,
		TotalPrice:  item.TotalPrice,
		NmID:        item.NmId,
		Brand:       item.Brand,
		Status:      item.Status,
	}
}

func ConvertOrderToDomain(order *models.Order, delivery *models.Delivery, payment *models.Payment, items []*models.Item) *domain.Order {
	return &domain.Order{
		OrderUID:          order.OrderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmID:              order.SmId,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,

		Delivery: ConvertDeliveryToDomain(delivery),

		Payment: ConvertPaymentToDomain(payment),

		Items: ConvertItemsToDomain(items),
	}
}

func ConvertDeliveryToDomain(delivery *models.Delivery) *domain.Delivery {
	return &domain.Delivery{
		Name:    delivery.Name,
		Phone:   delivery.Phone,
		Zip:     delivery.Zip,
		City:    delivery.City,
		Address: delivery.Address,
		Region:  delivery.Region,
		Email:   delivery.Email,
	}
}

func ConvertPaymentToDomain(payment *models.Payment) *domain.Payment {
	return &domain.Payment{
		Transaction:  payment.Transaction,
		RequestID:    payment.RequestId,
		Currency:     payment.Currency,
		Provider:     payment.Provider,
		Amount:       payment.Amount,
		PaymentDt:    payment.PaymentDt,
		Bank:         payment.Bank,
		DeliveryCost: payment.DeliveryCost,
		GoodsTotal:   payment.GoodsTotal,
		CustomFee:    payment.CustomFee,
	}
}

func ConvertItemsToModel(items []*domain.Item) []*models.Item {
	res := make([]*models.Item, len(items))
	for i, item := range items {
		res[i] = ConvertItemToModel(item)
	}

	return res
}

func ConvertItemToModel(item *domain.Item) *models.Item {
	return &models.Item{
		ChrtId:      item.ChrtID,
		TrackNumber: item.TrackNumber,
		Price:       item.Price,
		Rid:         item.Rid,
		Name:        item.Name,
		Sale:        item.Sale,
		Size:        item.Size,
		TotalPrice:  item.TotalPrice,
		NmId:        item.NmID,
		Brand:       item.Brand,
		Status:      item.Status,
	}
}

func ConvertOrderToModel(order *domain.Order) *models.Order {
	return &models.Order{
		OrderUid:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmId:              order.SmID,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}
}

func ConvertPaymentToModel(payment *domain.Payment) *models.Payment {
	return &models.Payment{
		Transaction:  payment.Transaction,
		RequestId:    payment.RequestID,
		Currency:     payment.Currency,
		Provider:     payment.Provider,
		Amount:       payment.Amount,
		PaymentDt:    payment.PaymentDt,
		Bank:         payment.Bank,
		DeliveryCost: payment.DeliveryCost,
		GoodsTotal:   payment.GoodsTotal,
		CustomFee:    payment.CustomFee,
	}
}

func ConvertDeliveryToModel(delivery *domain.Delivery) *models.Delivery {
	return &models.Delivery{
		Name:    delivery.Name,
		Phone:   delivery.Phone,
		Zip:     delivery.Zip,
		City:    delivery.City,
		Address: delivery.Address,
		Region:  delivery.Region,
		Email:   delivery.Email,
	}
}
