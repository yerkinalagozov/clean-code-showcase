package metrics

import (
	"context"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/service"
)

type orderRepoMetricsImpl struct {
	orderRepo *service.IOrderRepo
}

func NewOrderRepoMetrics(orderRepo *service.IOrderRepo) *orderRepoMetricsImpl {
	return &orderRepoMetricsImpl{orderRepo: orderRepo}
}

func (o *orderRepoMetricsImpl) Create(ctx context.Context, order ...entity.Order) ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepoMetricsImpl) Delete(ctx context.Context, orderID int) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepoMetricsImpl) GetBy(ctx context.Context, orderQuery entity.OrderQuery) (entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepoMetricsImpl) GetsBy(ctx context.Context, orderQuery entity.OrderQuery) ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}
