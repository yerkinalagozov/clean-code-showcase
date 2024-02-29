package service

import (
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type NewUserReq struct {
	id        int
	userName  string
	firstName string
	lastName  string
	email     string
	age       int
	isMarried string
	password  string
	status    string
}

func (n *NewUserReq) Id() int {
	return n.id
}

func (n *NewUserReq) SetId(id int) {
	n.id = id
}

func (n *NewUserReq) UserName() string {
	return n.userName
}

func (n *NewUserReq) SetUserName(userName string) {
	n.userName = userName
}

func (n *NewUserReq) FirstName() string {
	return n.firstName
}

func (n *NewUserReq) SetFirstName(firstName string) {
	n.firstName = firstName
}

func (n *NewUserReq) LastName() string {
	return n.lastName
}

func (n *NewUserReq) SetLastName(lastName string) {
	n.lastName = lastName
}

func (n *NewUserReq) Email() string {
	return n.email
}

func (n *NewUserReq) SetEmail(email string) {
	n.email = email
}

func (n *NewUserReq) Age() int {
	return n.age
}

func (n *NewUserReq) SetAge(age int) {
	n.age = age
}

func (n *NewUserReq) IsMarried() string {
	return n.isMarried
}

func (n *NewUserReq) SetIsMarried(isMarried string) {
	n.isMarried = isMarried
}

func (n *NewUserReq) Password() string {
	return n.password
}

func (n *NewUserReq) SetPassword(password string) {
	n.password = password
}

func (n *NewUserReq) Status() string {
	return n.status
}

func (n *NewUserReq) SetStatus(status string) {
	n.status = status
}

func (n *NewUserReq) MapToEntity() (entity.User, error) {
	var user entity.User
	user.SetID(n.id)
	user.SetUserName(n.userName)
	user.SetFirstLastName(n.firstName, n.lastName)
	err := user.SetEmail(n.email)
	if err != nil {
		return user, err
	}
	err = user.SetAge(n.age)
	if err != nil {
		return user, err
	}
	err = user.SetMarried(n.isMarried)
	if err != nil {
		return user, err
	}
	err = user.SetNewPassword(n.password)
	if err != nil {
		return user, err
	}
	err = user.SetStatus(n.status)
	if err != nil {
		return user, err
	}
	return user, nil
}

////////////// User //////////////

type UserReq struct {
	id        int
	userName  string
	firstName string
	lastName  string
	email     string
	age       int
	isMarried string
	password  string
	status    string
}

func (u *UserReq) MapToEntity() (entity.User, error) {
	var user entity.User
	user.SetID(u.id)
	user.SetUserName(u.userName)
	user.SetFirstLastName(u.firstName, u.lastName)
	err := user.SetEmail(u.email)
	if err != nil {
		return user, err
	}
	err = user.SetAge(u.age)
	if err != nil {
		return user, err
	}
	err = user.SetMarried(u.isMarried)
	if err != nil {
		return user, err
	}
	user.SetPassword(u.password)
	err = user.SetStatus(u.status)
	if err != nil {
		return user, err
	}
	return user, nil

}

//////////// Order /////////////

type NewOrderReq struct {
	id         int
	user       UserReq
	status     string
	orderItems []OrderItemReq
}

func (n *NewOrderReq) MapToEntity() (entity.Order, error) {
	var order entity.Order
	order.SetID(n.id)
	user, err := n.user.MapToEntity()
	if err != nil {
		return order, err
	}
	order.SetUser(user)
	err = order.SetOrderStatus(n.status)
	if err != nil {
		return order, err
	}
	var orderItems []entity.OrderItem
	for _, item := range n.orderItems {
		orderItem, err := item.MapToEntity()
		if err != nil {
			return order, err
		}
		orderItems = append(orderItems, orderItem)
	}
	order.SetOrderItems(orderItems)
	return order, nil

}

/////////// NewOrderItems /////////////

type NewOrderItemReq struct {
	id            int
	status        string
	product       ProductReq
	orderQuantity int
	priceAtOrder  float64
}

func (n *NewOrderItemReq) MapToEntity() (entity.OrderItem, error) {
	var orderItem entity.OrderItem
	orderItem.SetID(n.id)
	err := orderItem.SetOrderStatus(n.status)
	if err != nil {
		return orderItem, err
	}
	product, err := n.product.MapToEntity()
	if err != nil {
		return orderItem, err
	}
	orderItem.SetProduct(product)
	orderItem.SetQuantity(n.orderQuantity)
	orderItem.SetPriceAtOrder(n.priceAtOrder)
	return orderItem, nil

}

//////// OrderItem ////////////

type OrderItemReq struct {
	id            int
	status        string
	product       ProductReq
	orderQuantity int
	priceAtOrder  float64
}

func (o *OrderItemReq) MapToEntity() (entity.OrderItem, error) {
	var orderItem entity.OrderItem
	orderItem.SetID(o.id)
	err := orderItem.SetOrderStatus(o.status)
	if err != nil {
		return orderItem, err
	}
	product, err := o.product.MapToEntity()
	if err != nil {
		return orderItem, err
	}
	orderItem.SetProduct(product)
	orderItem.SetQuantity(o.orderQuantity)
	orderItem.SetPriceAtOrder(o.priceAtOrder)
	return orderItem, nil
}

//////////// NewProduct ///////////

type NewProductReq struct {
	id          int
	name        string
	description string
	tag         string
	price       float64
	quantity    int
}

func (p *NewProductReq) MapToEntity() (entity.ProductItems, error) {
	var product entity.ProductItems
	product.SetID(p.id)
	product.SetName(p.name)
	product.SetDescription(p.description)
	product.SetTag(p.tag)
	product.SetPrice(p.price)
	product.SetQuantity(p.quantity)
	return product, nil
}

//////////// Product ///////////

type ProductReq struct {
	id          int
	name        string
	description string
	tag         string
	price       float64
	quantity    int
}

func (p *ProductReq) MapToEntity() (entity.ProductItems, error) {
	var product entity.ProductItems
	product.SetID(p.id)
	product.SetName(p.name)
	product.SetDescription(p.description)
	product.SetTag(p.tag)
	product.SetPrice(p.price)
	product.SetQuantity(p.quantity)
	return product, nil
}
