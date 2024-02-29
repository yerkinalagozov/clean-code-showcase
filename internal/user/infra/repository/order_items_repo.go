package repository

import (
	"context"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/dbconnection"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type orderItemsRepo struct {
	pool dbconnection.Clients
}

func NewOrderItemsRepo(pool dbconnection.Clients) *orderItemsRepo {
	return &orderItemsRepo{pool: pool}
}

func (o orderItemsRepo) CreateItems(ctx context.Context, orderItem ...entity.OrderItem) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderItemsRepo) GetBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) (entity.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderItemsRepo) GetsBy(ctx context.Context, orderItemQuery entity.OrderItemQuery) ([]entity.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderItemsRepo) Update(ctx context.Context, orderItem entity.OrderItem, orderItemID int) error {
	//TODO implement me
	panic("implement me")
}

func (o orderItemsRepo) Delete(ctx context.Context, orderItemID int) error {
	//TODO implement me
	panic("implement me")
}
