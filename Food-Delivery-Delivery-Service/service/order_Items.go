package service

import (
	"context"

	"github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food/Food-Delivery-Delivery-Service/storage"
)

type OrderItemsService struct {
	storage storage.StorageI
	genproto.UnimplementedOrderItemServiceServer
}

func NewOrderItemsService(storage storage.StorageI) *OrderItemsService {
	return &OrderItemsService{storage: storage}
}

func (o *OrderItemsService) CreateOrderItem(ctx context.Context, req *genproto.CreateOrderItemRequest) (*genproto.OrderItemEmpty, error) {
	return o.storage.OrderItem().CreateOrderItem(req)
}

func (o *OrderItemsService) GetOrderItem(ctx context.Context, req *genproto.GetOrderItemRequest) (*genproto.OrderItem, error) {
	return o.storage.OrderItem().GetOrderItem(req)
}

func (o *OrderItemsService) UpdateOrderItem(ctx context.Context, req *genproto.UpdateOrderItemRequest) (*genproto.OrderItemEmpty, error) {
	return o.storage.OrderItem().UpdateOrderItem(req)
}

func (o *OrderItemsService) DeleteOrderItem(ctx context.Context, req *genproto.DeleteOrderItemRequest) (*genproto.OrderItemEmpty, error) {
	return o.storage.OrderItem().DeleteOrderItem(req)
}

func (o *OrderItemsService) ListOrderItems(ctx context.Context, req *genproto.GetAllOrderItemsRequest) (*genproto.OrderItemList, error) {
	return o.storage.OrderItem().ListOrderItems(req)
}
