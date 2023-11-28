package entity

import "time"

type OrderQueryOption func(*OrderQuery) error

type OrderQuery struct {
	id          []int
	user        []User
	orderStatus []OrderStatus
	orderItems  []OrderItem
	createAt    []time.Time
}

func NewOrderQuery(opts ...OrderQueryOption) (OrderQuery, error) {
	var q OrderQuery
	var err error
	for _, opt := range opts {
		err = opt(&q)
		if err != nil {
			return OrderQuery{}, err
		}
	}
	return q, nil
}

func OrderWithWhereById(id int) OrderQueryOption {
	return func(q *OrderQuery) error {
		q.id = append(q.id, id)
		return nil
	}
}

func OrderWithWhereByUserID(userID User) OrderQueryOption {
	return func(q *OrderQuery) error {
		q.user = append(q.user, userID)
		return nil
	}
}

func OrderWithWhereByOrderStatus(orderStatus OrderStatus) OrderQueryOption {
	return func(q *OrderQuery) error {
		q.orderStatus = append(q.orderStatus, orderStatus)
		return nil
	}
}

func OrderWithWhereByOrderItems(orderItems []OrderItem) OrderQueryOption {
	return func(q *OrderQuery) error {
		q.orderItems = append(q.orderItems, orderItems...)
		return nil
	}
}

func OrderWithWhereByCreateAt(createAt time.Time) OrderQueryOption {
	return func(q *OrderQuery) error {
		q.createAt = append(q.createAt, createAt)
		return nil
	}
}

func (o *OrderQuery) ID() []int {
	return o.id
}

func (o *OrderQuery) SetId(id []int) {
	o.id = id
}

func (o *OrderQuery) User() []User {
	return o.user
}

func (o *OrderQuery) SetUserID(userID []User) {
	o.user = userID
}

func (o *OrderQuery) OrderStatus() []string {
	var result []string
	for _, status := range o.orderStatus {
		result = append(result, status.String())
	}
	return result
}

func (o *OrderQuery) SetOrderStatus(orderStatus []OrderStatus) {
	o.orderStatus = orderStatus
}

func (o *OrderQuery) OrderItems() []OrderItem {
	return o.orderItems
}

func (o *OrderQuery) SetOrderItems(orderItems []OrderItem) {
	o.orderItems = orderItems
}

func (o *OrderQuery) CreateAt() []time.Time {
	return o.createAt
}

func (o *OrderQuery) SetCreateAt(createAt []time.Time) {
	o.createAt = createAt
}
