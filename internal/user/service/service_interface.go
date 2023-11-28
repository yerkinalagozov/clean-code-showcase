package service

import (
	"context"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type IProductRepo interface {
	Create(ctx context.Context, product ...entity.ProductItems) ([]entity.ProductItems, error)
	GetBy(ctx context.Context, productQuery entity.ProductQuery) (entity.ProductItems, error)
	GetsBy(ctx context.Context, productQuery entity.ProductQuery) ([]entity.ProductItems, error)
	Update(ctx context.Context, product ...entity.ProductItems) error
	Delete(ctx context.Context, productID int) error
}

type IUserRepo interface {
	Create(ctx context.Context, usersIn ...entity.User) ([]entity.User, error)
	GetBy(ctx context.Context, userQuery entity.UserQuery) (entity.User, error)
	GetsBy(ctx context.Context, userQuery entity.UserQuery) ([]entity.User, error)
	Update(ctx context.Context, user entity.User, userID int) error
	Delete(ctx context.Context, userID int) error
}

type IOrderRepo interface {
	Create(ctx context.Context, order ...entity.Order) ([]entity.Order, error)
	GetBy(ctx context.Context, orderQuery entity.OrderQuery) (entity.Order, error)
	GetsBy(ctx context.Context, orderQuery entity.OrderQuery) ([]entity.Order, error)
	Delete(ctx context.Context, orderID int) error
}

type IOrderItemRepo interface {
	Create(ctx context.Context, orderItem ...entity.OrderItem) (int, error)
	GetBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) (entity.OrderItem, error)
	GetsBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) ([]entity.OrderItem, error)
	Update(ctx context.Context, orderItem entity.OrderItem, orderItemID int) error
	Delete(ctx context.Context, orderItemID int) error
}

type IService interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUser(ctx context.Context, userRequest entity.UserRequest) (entity.User, error)
	GetUsers(ctx context.Context, userRequest entity.GetUsersRequest) ([]entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, userID int) error
	CreateProduct(ctx context.Context, product entity.ProductItems) (int, error)
	GetProduct(ctx context.Context, productRequest entity.GetProductRequest) (entity.ProductItems, error)
	GetProducts(ctx context.Context, productRequest entity.GetProductsRequest) ([]entity.ProductItems, error)
	UpdateProduct(ctx context.Context, product entity.ProductItems) error
	DeleteProduct(ctx context.Context, productID int) error
	CreateOrder(ctx context.Context, order entity.Order) (int, error)
	GetOrder(ctx context.Context, orderRequest entity.GetOrderRequest) (entity.Order, error)
	GetOrders(ctx context.Context, orderRequest entity.GetOrdersRequest) ([]entity.Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
	CreateOrderItem(ctx context.Context, orderItem entity.OrderItem) (int, error)
	GetOrderItem(ctx context.Context, orderItemRequest entity.GetOrderItemRequest) (entity.OrderItem, error)
	GetOrderItems(ctx context.Context, orderItemRequest entity.GetOrderItemsRequest) ([]entity.OrderItem, error)
	UpdateOrderItem(ctx context.Context, orderItem entity.OrderItem) error
	DeleteOrderItem(ctx context.Context, orderItemID int) error
}
