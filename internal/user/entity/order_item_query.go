package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItemQueryOption func(*OrderItemQuery) error

type OrderItemQuery struct {
	id            []int
	product       []ProductItems
	orderQuantity []int
	priceAtOrder  []decimal.Decimal
	createdAt     []time.Time
}

func NewOrderItemQuery(opts ...OrderItemQueryOption) (OrderItemQuery, error) {
	var q OrderItemQuery
	var err error
	for _, opt := range opts {
		err = opt(&q)
		if err != nil {
			return OrderItemQuery{}, err
		}
	}
	return q, nil
}

func OrderItemWithWhereById(id int) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.id = append(q.id, id)
		return nil
	}
}

func OrderItemWithWhereByIDs(id []int) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.id = append(q.id, id...)
		return nil
	}
}

func OrderItemWithWhereByProduct(product ProductItems) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.product = append(q.product, product)
		return nil
	}
}

func OrderItemWithWhereByOrderQuantity(orderQuantity int) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.orderQuantity = append(q.orderQuantity, orderQuantity)
		return nil
	}
}

func OrderItemWithWhereByPriceAtOrder(priceAtOrder decimal.Decimal) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.priceAtOrder = append(q.priceAtOrder, priceAtOrder)
		return nil
	}
}

func OrderItemWithWhereByCreatedAt(createdAt time.Time) OrderItemQueryOption {
	return func(q *OrderItemQuery) error {
		q.createdAt = append(q.createdAt, createdAt)
		return nil
	}
}

func (o *OrderItemQuery) Id() []int {
	return o.id
}

func (o *OrderItemQuery) SetId(id []int) {
	o.id = id
}

func (o *OrderItemQuery) Product() []ProductItems {
	return o.product
}

func (o *OrderItemQuery) SetProduct(product []ProductItems) {
	o.product = product
}

func (o *OrderItemQuery) OrderQuantity() []int {
	return o.orderQuantity
}

func (o *OrderItemQuery) SetOrderQuantity(orderQuantity []int) {
	o.orderQuantity = orderQuantity
}

func (o *OrderItemQuery) PriceAtOrder() []decimal.Decimal {
	return o.priceAtOrder
}

func (o *OrderItemQuery) SetPriceAtOrder(priceAtOrder []decimal.Decimal) {
	o.priceAtOrder = priceAtOrder
}

func (o *OrderItemQuery) CreatedAt() []time.Time {
	return o.createdAt
}

func (o *OrderItemQuery) SetCreatedAt(createdAt []time.Time) {
	o.createdAt = createdAt
}
