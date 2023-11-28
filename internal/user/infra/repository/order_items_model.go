package repository

import (
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type OrderItemsData struct {
	ID           int         `json:"id"`
	Status       string      `json:"order_status"`
	Product      ProductData `json:"product"`
	Quantity     int         `json:"quantity"`
	PriceAtOrder float64     `json:"price_at_order"`
	CreatedAt    time.Time   `json:"created_at"`
}

func (o *OrderItemsData) MapToRepoNew(orderItemsData entity.OrderItem) {
	o.ID = orderItemsData.ID()
	o.Status = orderItemsData.Status().String()
	o.Product.MapToRepo(orderItemsData.Product())
	o.Quantity = orderItemsData.Quantity()
	var ok bool
	o.PriceAtOrder, ok = orderItemsData.PriceAtOrder().Float64()
	if !ok {
		o.PriceAtOrder = 0
	}
	o.CreatedAt = time.Now().UTC()
}

func (o *OrderItemsData) MapToRepoNewList(orderItemsDataIn []entity.OrderItem) []OrderItemsData {
	var result []OrderItemsData
	for _, orderItem := range orderItemsDataIn {
		orderItemsData := OrderItemsData{}
		orderItemsData.MapToRepoNew(orderItem)
		result = append(result, orderItemsData)
	}
	return result
}

func (o *OrderItemsData) MapToRepo(orderItemsData entity.OrderItem) {
	o.ID = orderItemsData.ID()
	o.Product.MapToRepo(orderItemsData.Product())
	o.Quantity = orderItemsData.Quantity()
	var ok bool
	o.PriceAtOrder, ok = orderItemsData.PriceAtOrder().Float64()
	if !ok {
		o.PriceAtOrder = 0
	}
}

func (o *OrderItemsData) MapToRepoList(orderItemsDataIn []entity.OrderItem) []OrderItemsData {
	var result []OrderItemsData
	for _, orderItem := range orderItemsDataIn {
		orderItemsData := OrderItemsData{}
		orderItemsData.MapToRepo(orderItem)
		result = append(result, orderItemsData)
	}
	return result
}

func (o *OrderItemsData) MapToEntity() (entity.OrderItem, error) {
	var orderItem entity.OrderItem
	orderItem.SetID(o.ID)
	orderItem.SetStatus(entity.OrderItemStatus(o.Status))
	orderItem.SetProduct(o.Product.MapToEntity())
	orderItem.SetQuantity(o.Quantity)
	orderItem.SetPriceAtOrder(o.PriceAtOrder)
	orderItem.SetCreatedAt(o.CreatedAt)
	return orderItem, nil
}

func (o *OrderItemsData) MapToEntityList(orderItemsDataIn []OrderItemsData) ([]entity.OrderItem, error) {
	var result []entity.OrderItem
	for _, orderItem := range orderItemsDataIn {
		entityOrderItemTemp, err := orderItem.MapToEntity()
		if err != nil {
			return nil, err
		}
		result = append(result, entityOrderItemTemp)
	}
	return result, nil
}
