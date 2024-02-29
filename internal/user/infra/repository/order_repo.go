package repository

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/dbconnection"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type orderRepoImpl struct {
	pool dbconnection.Clients
}

func NewOrderRepo(pool dbconnection.Clients) *orderRepoImpl {
	return &orderRepoImpl{pool: pool}
}

func (o *orderRepoImpl) Create(ctx context.Context, order ...entity.Order) ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepoImpl) DeleteOrder(ctx context.Context, orderID int) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepoImpl) CreateOrder(ctx context.Context, ordersIn ...entity.Order) ([]entity.Order, error) {
	var orderData OrderData
	ordersDataList := orderData.MapToRepoNewList(ordersIn)
	query := `INSERT INTO orders (user, status, order_items, created_at) VALUES ($1, $2, $3) RETURNING id`

	batch := &pgx.Batch{}

	for _, order := range ordersDataList {
		batch.Queue(query, order.User, order.Status, order.OrderItems, order.CreatedAt)
	}
	br := o.pool.SendBatch(ctx, batch)
	defer func() {
		err := br.Close()
		if err != nil {
			slog.Error("orderRepo.CreateOrder.br.Close, error while closing batch results", slog.Any("error", err))
		}
	}()

	var createdOrders []OrderData
	for _, order := range ordersDataList {
		var id int
		err := br.QueryRow().Scan(&id)
		if err != nil {
			return nil, err
		}
		order.ID = id
		createdOrders = append(createdOrders, order)
	}
	var resultOrder OrderData
	return resultOrder.MapToEntityList(createdOrders)
}

func (o *orderRepoImpl) GetByOrder(ctx context.Context, orderQuery entity.OrderQuery) (entity.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("user", "status", "order_items", "created_at").From("orders")
	sqb = o.generateUpdateWhere(sqb, orderQuery)

	query, args, err := sqb.ToSql()
	if err != nil {
		return entity.Order{}, commonentity.NewDatabaseError(
			errors.Wrapf(err, "orderRepoImpl.GetBy.ToSql, error while getting orders"),
			commonentity.ErrUnknownStatus,
			"error while getting orders",
		)
	}
	row := o.pool.QueryRow(ctx, query, args...)
	var order OrderData
	err = row.Scan(&order.User, &order.Status, &order.OrderItems, &order.CreatedAt)
	if err != nil {
		return entity.Order{}, commonentity.NewDatabaseError(
			errors.Wrapf(err, "orderRepoImpl.GetBy.Scan, error while getting orders"),
			commonentity.ErrUnknownStatus,
			"error while getting orders",
		)
	}
	return order.MapToEntity()
}

func (o *orderRepoImpl) GetsByOrder(ctx context.Context, orderQuery entity.OrderQuery) ([]entity.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqb := psql.Select("user", "status", "order_items", "created_at").From("orders")
	sqb = o.generateUpdateWhere(sqb, orderQuery)

	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "orderRepoImpl.GetsBy.ToSql, error while getting orders"),
			commonentity.ErrUnknownStatus,
			"error while getting orders",
		)
	}
	rows, err := o.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, commonentity.NewDatabaseError(
			errors.Wrapf(err, "orderRepoImpl.GetsBy.Query, error while getting orders"),
			commonentity.ErrUnknownStatus,
			"error while getting orders",
		)
	}
	defer rows.Close()

	var orders []OrderData
	for rows.Next() {
		var order OrderData
		err = rows.Scan(&order.User, &order.Status, &order.OrderItems, &order.CreatedAt)
		if err != nil {
			return nil, commonentity.NewDatabaseError(
				errors.Wrapf(err, "orderRepoImpl.GetsBy.Scan, error while getting orders"),
				commonentity.ErrUnknownStatus,
				"error while getting orders",
			)
		}
		orders = append(orders, order)
	}
	var resultOrder OrderData
	return resultOrder.MapToEntityList(orders)
}

func (o *orderRepoImpl) Update(ctx context.Context, order ...entity.Order) (entity.Order, error) {
	return entity.Order{}, commonentity.NewDatabaseError()
}

func (o *orderRepoImpl) generateUpdateWhere(sqb sq.SelectBuilder, orderQuery entity.OrderQuery) sq.SelectBuilder {
	if len(orderQuery.ID()) > 0 {
		sqb = sqb.Where(sq.Eq{"id": orderQuery.ID})
	}
	if len(orderQuery.User()) > 0 {
		sqb = sqb.Where(sq.Eq{"user": orderQuery.User})
	}
	if len(orderQuery.OrderStatus()) > 0 {
		sqb = sqb.Where(sq.Eq{"status": orderQuery.OrderStatus()})
	}
	if len(orderQuery.OrderItems()) > 0 {
		sqb = sqb.Where(sq.Eq{"order_items": orderQuery.OrderItems()})
	}
	if len(orderQuery.CreateAt()) > 0 {
		sqb = sqb.Where(sq.Eq{"created_at": orderQuery.CreateAt()})
	}
	return sqb
}
