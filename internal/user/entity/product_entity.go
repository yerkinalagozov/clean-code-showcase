package entity

import (
	"time"
)

type ProductItems struct {
	id          int
	name        string
	description string
	tag         string
	price       float64
	quantity    int
	createdAt   time.Time
}

func (p *ProductItems) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProductItems) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *ProductItems) Price() float64 {
	return p.price
}

func (p *ProductItems) SetPrice(price float64) {
	p.price = price
}

func (p *ProductItems) Name() string {
	return p.name
}

func (p *ProductItems) SetName(name string) {
	p.name = name
}

func (p *ProductItems) IsZero() bool {
	return p.quantity == 0
}

func (p *ProductItems) IsEmpty() bool {
	return p.id == 0
}

func (p *ProductItems) ID() int {
	return p.id
}

func (p *ProductItems) SetID(id int) {
	p.id = id
}

func (p *ProductItems) Description() string {
	return p.description
}

func (p *ProductItems) SetDescription(description string) {
	p.description = description
}

func (p *ProductItems) Tag() string {
	return p.tag
}

func (p *ProductItems) SetTag(tag string) {
	p.tag = tag
}

func (p *ProductItems) Quantity() int {
	return p.quantity
}

func (p *ProductItems) SetQuantity(quantity int) {
	p.quantity = quantity
}
