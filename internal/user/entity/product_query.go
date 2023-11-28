package entity

import "time"

type ProductQueryOption func(*ProductQuery) error

type ProductQuery struct {
	id          []int
	description []string
	tag         []string
	quantity    []int
	createdAt   []time.Time
}

func (p *ProductQuery) CreatedAt() []time.Time {
	return p.createdAt
}

func (p *ProductQuery) SetCreatedAt(createdAt []time.Time) {
	p.createdAt = createdAt
}

func NewProductQuery(opts ...ProductQueryOption) (ProductQuery, error) {
	var q ProductQuery
	var err error
	for _, opt := range opts {
		err = opt(&q)
		if err != nil {
			return ProductQuery{}, err
		}
	}
	return q, nil
}

func ProductWithWhereById(id int) ProductQueryOption {
	return func(q *ProductQuery) error {
		q.id = append(q.id, id)
		return nil
	}
}

func ProductWithWhereByDescription(description string) ProductQueryOption {
	return func(q *ProductQuery) error {
		q.description = append(q.description, description)
		return nil
	}
}

func ProductWithWhereByTag(tag string) ProductQueryOption {
	return func(q *ProductQuery) error {
		q.tag = append(q.tag, tag)
		return nil
	}
}

func ProductWithWhereByQuantity(quantity int) ProductQueryOption {
	return func(q *ProductQuery) error {
		q.quantity = append(q.quantity, quantity)
		return nil
	}
}

func ProductWithWhereByCreatedAt(createdAt time.Time) ProductQueryOption {
	return func(q *ProductQuery) error {
		q.createdAt = append(q.createdAt, createdAt)
		return nil
	}
}

func (p *ProductQuery) IDs() []int {
	return p.id
}

func (p *ProductQuery) SetID(id []int) {
	p.id = id
}

func (p *ProductQuery) Description() []string {
	return p.description
}

func (p *ProductQuery) SetDescription(description []string) {
	p.description = description
}

func (p *ProductQuery) Tag() []string {
	return p.tag
}

func (p *ProductQuery) SetTag(tag []string) {
	p.tag = tag
}

func (p *ProductQuery) Quantity() []int {
	return p.quantity
}

func (p *ProductQuery) SetQuantity(quantity []int) {
	p.quantity = quantity
}
