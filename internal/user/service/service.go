package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type service struct {
	userRepo      IUserRepo
	productRepo   IProductRepo
	orderRepo     IOrderRepo
	orderItemRepo IOrderItemRepo
}

func NewService(userRepo IUserRepo, productRepo IProductRepo, orderRepo IOrderRepo, orderItemRepo IOrderItemRepo) *service {
	return &service{
		userRepo:      userRepo,
		productRepo:   productRepo,
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (s *service) NewUser(ctx context.Context, user NewUserReq) (int, error) {
	newUser, err := user.MapToEntity()
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewUser.MapToEntity, failed to map user to entity")
	}
	newUserID, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewUser.Create, failed to create user")
	}
	return newUserID[0].ID(), nil
}

func (s *service) NewOrder(ctx context.Context, orderIn NewOrderReq) (int, error) {
	newOrder, err := orderIn.MapToEntity()
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewOrder.MapToEntity, failed to map order to entity")
	}
	existOrder, err := s.getOrder(ctx, newOrder)
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewOrder.GetOrder, failed to get order")
	}
	if !existOrder.IsEmpty() {
		return 0, errors.WithMessagef(err, "service.NewOrder, failed to check order, order %d already exist", existOrder.ID())
	}
	// Check user exist
	existUser, err := s.getUser(ctx, newOrder.User())
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewOrder.GetUser, failed to get user")
	}
	// set user
	newOrder.SetUser(existUser)

	// Check orderItem exist
	existOrderItems, err := s.getOrdersItem(ctx, newOrder.OrderItems()...)
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewOrder.GetOrderItem, failed to get orderItem")
	}
	// check len orderItem and len existOrderItem
	if len(existOrderItems) != len(newOrder.OrderItems()) {
		return 0, errors.WithMessagef(err, "service.NewOrder, failed to check orderItem, orderItem %v already exist", existOrderItems)
	}
	// set orderItem
	newOrder.SetOrderItems(existOrderItems)
	// Create order
	orderID, err := s.orderRepo.Create(ctx, newOrder)
	if err != nil {
		return 0, errors.WithMessagef(err, "service.NewOrder.Create, failed to create order")
	}
	return orderID[0].ID(), nil
}

func (s *service) NewOrderItems(ctx context.Context, orderItemsIn NewOrderItemReq) (entity.OrderItem, error) {
	newOrderItem, err := orderItemsIn.MapToEntity()
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.NewOrderItems.MapToEntity, failed to map orderItem to entity")
	}
	orderItems, err := s.getOrderItem(ctx, newOrderItem)
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.NewOrderItems.GetOrderItem, failed to get orderItem")
	}
	// Check product exist
	existProduct, err := s.getProduct(ctx, orderItems.Product())
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.NewOrderItems.GetProduct, failed to get product")
	}
	orderItems.SetProduct(existProduct)

	// Create orderItem
	orderItemID, err := s.orderItemRepo.Create(ctx, orderItems)
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.NewOrderItems.Create, failed to create orderItem")
	}
	orderItems.SetID(orderItemID)
	return orderItems, nil
}

func (s *service) getOrderItem(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error) {
	orderItemQuery, err := entity.NewOrderItemQuery(entity.OrderItemWithWhereById(orderItem.ID()))
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.GetOrderItem.NewOrderItemQuery, failed to create orderItemQuery")
	}
	existOrderItem, err := s.orderItemRepo.GetBy(ctx, orderItemQuery)
	if err != nil {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.GetOrderItem.GetBy, failed to get orderItem")
	}
	if existOrderItem.IsEmpty() {
		return entity.OrderItem{}, errors.WithMessagef(err, "service.GetOrderItem.IsEmpty, failed to check orderItem, orderItem %d not exist", existOrderItem.ID())
	}
	return existOrderItem, nil
}

func (s *service) getOrdersItem(ctx context.Context, orderItem ...entity.OrderItem) ([]entity.OrderItem, error) {
	orderItemsID := make([]int, len(orderItem))
	for _, item := range orderItem {
		orderItemsID = append(orderItemsID, item.ID())
	}
	orderItemQuery, err := entity.NewOrderItemQuery(entity.OrderItemWithWhereByIDs(orderItemsID))
	if err != nil {
		return nil, errors.WithMessagef(err, "service.GetOrderItem.NewOrderItemQuery, failed to create orderItemQuery")
	}
	existOrderItem, err := s.orderItemRepo.GetsBy(ctx, orderItemQuery)
	if err != nil {
		return nil, errors.WithMessagef(err, "service.GetOrderItem.GetBy, failed to get orderItem")
	}
	if len(existOrderItem) == 0 {
		return nil, errors.WithMessagef(err, "service.GetOrderItem.IsEmpty, failed to check orderItem, orderItem %v not exist", orderItemsID)
	}
	return existOrderItem, nil
}

func (s *service) getProduct(ctx context.Context, product entity.ProductItems) (entity.ProductItems, error) {
	productQuery, err := entity.NewProductQuery(entity.ProductWithWhereById(product.ID()))
	if err != nil {
		return entity.ProductItems{}, errors.WithMessagef(err, "service.GetProduct.NewProductQuery, failed to create productQuery")
	}
	existProduct, err := s.productRepo.GetBy(ctx, productQuery)
	if err != nil {
		return entity.ProductItems{}, errors.WithMessagef(err, "service.GetProduct.GetBy, failed to get product")
	}
	if existProduct.IsEmpty() {
		return entity.ProductItems{}, errors.WithMessagef(err, "service.GetProduct.IsEmpty, failed to check product, product %d not exist", existProduct.ID())
	}
	return existProduct, nil
}

func (s *service) getOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	orderQuery, err := entity.NewOrderQuery(entity.OrderWithWhereById(order.ID()))
	if err != nil {
		return entity.Order{}, errors.WithMessagef(err, "service.GetOrder.NewOrderQuery, failed to create orderQuery")
	}
	existOrder, err := s.orderRepo.GetBy(ctx, orderQuery)
	if err != nil {
		return entity.Order{}, errors.WithMessagef(err, "service.GetOrder.GetBy, failed to get order")
	}
	if existOrder.IsEmpty() {
		return entity.Order{}, errors.WithMessagef(err, "service.GetOrder.IsEmpty, failed to check order, order %d not exist", existOrder.ID())
	}
	return existOrder, nil
}

func (s *service) getUser(ctx context.Context, user entity.User) (entity.User, error) {
	userQuery, err := entity.NewUserQuery(entity.UserWithWhereById(user.ID()))
	if err != nil {
		return entity.User{}, errors.WithMessagef(err, "service.GetUser.NewUserQuery, failed to create userQuery")
	}
	existUser, err := s.userRepo.GetBy(ctx, userQuery)
	if err != nil {
		return entity.User{}, errors.WithMessagef(err, "service.GetUser.GetBy, failed to get user")
	}
	if existUser.IsEmpty() {
		return entity.User{}, errors.WithMessagef(err, "service.GetUser.IsEmpty, failed to check user, user %d not exist", existUser.ID())
	}
	return existUser, nil
}
