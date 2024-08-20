package service

import (
	"context"

	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/storage"
)

type CartService struct {
	storage storage.StorageI
	genproto.UnimplementedCartServiceServer
}

func NewCartService(storage storage.StorageI) *CartService {
	return &CartService{storage: storage}
}

func (c *CartService) CreateCart(ctx context.Context, req *genproto.CreateCartReq) (*genproto.CartEmpty, error) {
	return c.storage.Cart().CreateCart(req)
}

func (c *CartService) GetCart(ctx context.Context, req *genproto.GetByIdCartRequest) (*genproto.Cart, error) {
	return c.storage.Cart().GetCart(req)
}

func (c *CartService) GetAllCarts(ctx context.Context, req *genproto.GetAllCartsReq) (*genproto.GetAllCartsRes, error) {
	return c.storage.Cart().GetAllCarts(req)
}

func (c *CartService) UpdateCart(ctx context.Context, req *genproto.UpdateCartReq) (*genproto.UpdateCartRes, error) {
	return c.storage.Cart().UpdateCart(req)
}

func (c *CartService) DeleteCart(ctx context.Context, req *genproto.DeleteCartRequest) (*genproto.DeleteCartResp, error) {
	return c.storage.Cart().DeleteCart(req)
}
