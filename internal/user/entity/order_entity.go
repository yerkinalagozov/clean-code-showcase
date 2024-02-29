package entity

import (
	"time"
)

type OrderStatus string

func (o *OrderStatus) String() string {
	return string(*o)
}

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusCanceled  OrderStatus = "canceled"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusRefunded  OrderStatus = "refunded"
)

type Order struct {
	id         int
	user       User
	status     OrderStatus
	orderItems []OrderItem
	createdAt  time.Time
}

func (o *Order) User() User {
	return o.user
}

func (o *Order) SetUser(user User) {
	o.user = user
}

func (o *Order) IsEmpty() bool {
	return o.id == 0
}

func (o *Order) ID() int {
	return o.id
}

func (o *Order) SetID(id int) {
	o.id = id
}

func (o *Order) UserID() int {
	return o.user.id
}

func (o *Order) OrderStatus() *OrderStatus {
	return &o.status
}

//func (o *Order) SetOrderStatus(orderStatus string) error {
//	var err error
//	switch orderStatus {
//	case string(OrderStatusCreated):
//		o.status = OrderStatusCreated
//	case string(OrderStatusCanceled):
//		o.status = OrderStatusCanceled
//	case string(OrderStatusPaid):
//		o.status = OrderStatusPaid
//	case string(OrderStatusDelivered):
//		o.status = OrderStatusDelivered
//	case string(OrderStatusRefunded):
//		o.status = OrderStatusRefunded
//	default:
//		err = commonentity.ErrOrderStatusIsNotValid
//	}
//	return err
//}

func (o *Order) OrderItems() []OrderItem {
	return o.orderItems
}

func (o *Order) SetOrderItems(orderItems []OrderItem) {
	o.orderItems = orderItems
}

func (o *Order) CreateAt() time.Time {
	return o.createdAt
}

func (o *Order) SetCreateAt(createAt time.Time) {
	o.createdAt = createAt
}
