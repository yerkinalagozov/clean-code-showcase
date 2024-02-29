package repository

import (
	"fmt"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type OrderData struct {
	ID         int              `json:"id"`
	User       UserData         `json:"user_id"`
	Status     string           `json:"order_status"`
	OrderItems []OrderItemsData `json:"order_ids"`
	CreatedAt  time.Time        `json:"created_at"`
}

func (o *OrderData) MapToRepoNew(order entity.Order) {
	o.ID = order.ID()
	o.User.MapToRepo(order.User())
	o.Status = order.OrderStatus().String()
	var orderItemsData OrderItemsData
	o.OrderItems = orderItemsData.MapToRepoNewList(order.OrderItems())
	o.CreatedAt = time.Now().UTC()
}

func (o *OrderData) MapToRepoNewList(orders []entity.Order) []OrderData {
	var result []OrderData
	for _, order := range orders {
		orderData := OrderData{}
		orderData.MapToRepoNew(order)
		result = append(result, orderData)
	}
	return result
}

func (o *OrderData) MapToRepo(order entity.Order) {
	o.ID = order.ID()
	o.User.MapToRepo(order.User())
	o.Status = order.OrderStatus().String()
	var orderItemsData OrderItemsData
	o.OrderItems = orderItemsData.MapToRepoList(order.OrderItems())
}

func (o *OrderData) MapToRepoList(orders []entity.Order) []OrderData {
	var result []OrderData
	for _, order := range orders {
		orderData := OrderData{}
		orderData.MapToRepo(order)
		result = append(result, orderData)
	}
	return result
}

func (o *OrderData) MapToEntity() (entity.Order, error) {
	var order entity.Order
	order.SetID(o.ID)
	entityUser, err := o.User.MapToEntity()
	if err != nil {
		return entity.Order{}, err
	}
	order.SetUser(entityUser)
	err = order.SetOrderStatus(o.Status)
	if err != nil {
		return entity.Order{}, fmt.Errorf("order status: %w", err)
	}
	var orderItemsData OrderItemsData
	entityOrderItems, err := orderItemsData.MapToEntityList(o.OrderItems)
	if err != nil {
		return entity.Order{}, fmt.Errorf("order items: %w", err)
	}
	order.SetOrderItems(entityOrderItems)
	order.SetCreateAt(o.CreatedAt)
	return order, nil
}

func (o *OrderData) MapToEntityList(orders []OrderData) ([]entity.Order, error) {
	var result []entity.Order
	for _, order := range orders {
		entityOrderTemp, err := order.MapToEntity()
		if err != nil {
			return nil, err
		}
		result = append(result, entityOrderTemp)
	}
	return result, nil
}
