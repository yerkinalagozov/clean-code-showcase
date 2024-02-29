package entity

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
)

type OrderItemStatus string

func (o *OrderItemStatus) String() string {
	return o.String()
}

const (
	OrderItemStatusCreated   OrderItemStatus = "created"
	OrderItemStatusCanceled  OrderItemStatus = "canceled"
	OrderItemStatusPaid      OrderItemStatus = "paid"
	OrderItemStatusDelivered OrderItemStatus = "delivered"
	OrderItemStatusRefunded  OrderItemStatus = "refunded"
)

type OrderItem struct {
	id            int
	status        OrderItemStatus
	product       ProductItems
	orderQuantity int
	priceAtOrder  decimal.Decimal
	createdAt     time.Time
}

func (o *OrderItem) Status() *OrderItemStatus {
	return &o.status
}

func (o *OrderItem) SetOrderStatus(orderStatus string) error {
	var err error
	switch orderStatus {
	case string(OrderItemStatusCreated):
		o.status = OrderItemStatusCreated
	case string(OrderItemStatusCanceled):
		o.status = OrderItemStatusCanceled
	case string(OrderItemStatusPaid):
		o.status = OrderItemStatusPaid
	case string(OrderItemStatusDelivered):
		o.status = OrderItemStatusDelivered
	case string(OrderItemStatusRefunded):
		o.status = OrderItemStatusRefunded
	default:
		err = commonentity.ErrOrderItemStatusIsNotValid
	}
	return err
}

func (o *OrderItem) Product() ProductItems {
	return o.product
}

func (o *OrderItem) SetProduct(product ProductItems) {
	o.product = product
}

func (o *OrderItem) IsEmpty() bool {
	return o.id == 0
}

func (o *OrderItem) ID() int {
	return o.id
}

func (o *OrderItem) SetID(orderID int) {
	o.id = orderID
}

func (o *OrderItem) ProductID() int {
	return o.product.id
}

func (o *OrderItem) SetProductID(productID int) {
	o.product.id = productID
}

func (o *OrderItem) Quantity() int {
	return o.orderQuantity
}

func (o *OrderItem) SetQuantity(quantity int) {
	o.orderQuantity = quantity
}

func (o *OrderItem) PriceAtOrder() decimal.Decimal {
	return o.priceAtOrder
}

func (o *OrderItem) SetPriceAtOrder(priceAtOrder float64) {
	o.priceAtOrder = decimal.NewFromFloat(priceAtOrder)
}

func (o *OrderItem) CreatedAt() time.Time {
	return o.createdAt
}

func (o *OrderItem) SetCreatedAt(createdAt time.Time) {
	o.createdAt = createdAt
}
