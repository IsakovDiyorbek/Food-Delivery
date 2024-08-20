package service

import (
	"context"

	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/storage"

	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
)

type OrderService struct {
	storage storage.StorageI
	genproto.UnimplementedOrderServiceServer
}

func NewOrderService(storage storage.StorageI) *OrderService {
	return &OrderService{storage: storage}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *genproto.CreateOrderRequest) (*genproto.OrderEmpty, error) {
	return o.storage.Order().CreateOrder(req)
}

func (o *OrderService) GetOrder(ctx context.Context, req *genproto.GetOrderRequest) (*genproto.Order, error) {
	return o.storage.Order().GetOrder(req)
}

func (o *OrderService) UpdateOrder(ctx context.Context, req *genproto.UpdateOrderRequest) (*genproto.OrderEmpty, error) {
	return o.storage.Order().UpdateOrder(req)
}

func (o *OrderService) DeleteOrder(ctx context.Context, req *genproto.DeleteOrderRequest) (*genproto.OrderEmpty, error) {
	return o.storage.Order().DeleteOrder(req)
}

func (o *OrderService) ListOrders(ctx context.Context, req *genproto.GetAllOrdersRequest) (*genproto.OrderList, error) {
	return o.storage.Order().ListOrders(req)
}
