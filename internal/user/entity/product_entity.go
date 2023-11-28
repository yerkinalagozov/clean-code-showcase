package entity

import (
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
)

type ProductItems struct {
	id          int
	name        commonentity.CustomString
	description commonentity.CustomString
	tag         commonentity.CustomString
	price       commonentity.CustomFloat
	quantity    commonentity.CustomInt
	createdAt   time.Time
}

func (p *ProductItems) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProductItems) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *ProductItems) Price() commonentity.CustomFloat {
	return p.price
}

func (p *ProductItems) SetPrice(price commonentity.CustomFloat) {
	p.price = price
}

func (p *ProductItems) Name() commonentity.CustomString {
	return p.name
}

func (p *ProductItems) SetName(name commonentity.CustomString) {
	p.name = name
}

func (p *ProductItems) IsZero() bool {
	return p.quantity.Val == 0
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

func (p *ProductItems) Description() commonentity.CustomString {
	return p.description
}

func (p *ProductItems) SetDescription(description commonentity.CustomString) {
	p.description = description
}

func (p *ProductItems) Tag() commonentity.CustomString {
	return p.tag
}

func (p *ProductItems) SetTag(tag commonentity.CustomString) {
	p.tag = tag
}

func (p *ProductItems) Quantity() commonentity.CustomInt {
	return p.quantity
}

func (p *ProductItems) SetQuantity(quantity commonentity.CustomInt) {
	p.quantity = quantity
}
