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
	Update(ctx context.Context, usersIn ...entity.User) error
	Delete(ctx context.Context, userID int) error
}

type IOrderRepo interface {
	CreateOrder(ctx context.Context, order ...entity.Order) ([]entity.Order, error)
	GetByOrder(ctx context.Context, orderQuery entity.OrderQuery) (entity.Order, error)
	GetsByOrder(ctx context.Context, orderQuery entity.OrderQuery) ([]entity.Order, error)
}

type IOrderItemRepo interface {
	CreateItems(ctx context.Context, orderItem ...entity.OrderItem) (int, error)
	GetBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) (entity.OrderItem, error)
	GetsBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) ([]entity.OrderItem, error)
	Update(ctx context.Context, orderItem entity.OrderItem, orderItemID int) error
	Delete(ctx context.Context, orderItemID int) error
}

type IService interface {
	NewUser(ctx context.Context, user NewUserReq) (int, error)
	NewOrder(ctx context.Context, orderIn NewOrderReq) (int, error)
	NewOrderItems(ctx context.Context, orderItemsIn NewOrderItemReq) (entity.OrderItem, error)
	//GetUser(ctx context.Context, userRequest UserRequest) (entity.User, error)
	//GetUsers(ctx context.Context, userRequest GetUsersRequest) ([]entity.User, error)
	//UpdateUser(ctx context.Context, user entity.User) error
	//DeleteUser(ctx context.Context, userID int) error
	//CreateProduct(ctx context.Context, product entity.ProductItems) (int, error)
	//GetProduct(ctx context.Context, productRequest GetProductRequest) (entity.ProductItems, error)
	//GetProducts(ctx context.Context, productRequest GetProductsRequest) ([]entity.ProductItems, error)
	//UpdateProduct(ctx context.Context, product entity.ProductItems) error
	//DeleteProduct(ctx context.Context, productID int) error
	//CreateOrder(ctx context.Context, order entity.Order) (int, error)
	//GetOrder(ctx context.Context, orderRequest GetOrderRequest) (entity.Order, error)
	//GetOrders(ctx context.Context, orderRequest GetOrdersRequest) ([]entity.Order, error)
	//DeleteOrder(ctx context.Context, orderID int) error
	//CreateOrderItem(ctx context.Context, orderItem entity.OrderItem) (int, error)
	//GetOrderItem(ctx context.Context, orderItemRequest GetOrderItemRequest) (entity.OrderItem, error)
	//GetOrderItems(ctx context.Context, orderItemRequest GetOrderItemsRequest) ([]entity.OrderItem, error)
	//UpdateOrderItem(ctx context.Context, orderItem entity.OrderItem) error
	//DeleteOrderItem(ctx context.Context, orderItemID int) error
}

type IServiceAdmin interface {
	NewUser(ctx context.Context, user NewUserReq) (int, error)
}
