package service

import (
	"context"
	"github.com/Food/Food-Delivery-Delivery-Service/storage"

	"github.com/Food/Food-Delivery-Delivery-Service/genproto"
)

type ProductService struct {
	storage storage.StorageI
	genproto.UnimplementedProductServiceServer
}

func NewProductService(storage storage.StorageI) *ProductService {
	return &ProductService{storage: storage}
}

func (p *ProductService) CreateProduct(ctx context.Context, req *genproto.CreateProductRequest) (*genproto.ProductEmpty, error) {
	return p.storage.Product().CreateProduct(req)
}

func (p *ProductService) GetProduct(ctx context.Context, req *genproto.GetProductRequest) (*genproto.Product, error) {
	return p.storage.Product().GetProduct(req)
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *genproto.UpdateProductRequest) (*genproto.ProductEmpty, error) {
	return p.storage.Product().UpdateProduct(req)
}

func (p *ProductService) DeleteProduct(ctx context.Context, req *genproto.DeleteProductRequest) (*genproto.ProductEmpty, error) {
	return p.storage.Product().DeleteProduct(req)
}

func (p *ProductService) ListProducts(ctx context.Context, req *genproto.GetAllProductsRequest) (*genproto.ProductList, error) {
	return p.storage.Product().ListProducts(req)
}