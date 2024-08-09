package service

import (
	"context"

	"github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food/Food-Delivery-Delivery-Service/storage"
)

type CourierLocationService struct {
	storage storage.StorageI
	genproto.UnimplementedCourierLocationServiceServer
}

func NewCourierLocationService(storage storage.StorageI) *CourierLocationService {
	return &CourierLocationService{storage: storage}
}

func (c *CourierLocationService) CreateCourierLocation(ctx context.Context, req *genproto.CreateCourierLocationRequest) (*genproto.Empty, error) {
	return c.storage.CourierLocation().CreateCourierLocation(req)
}

func (c *CourierLocationService) GetCourierLocation(ctx context.Context, req *genproto.GetCourierLocationRequest) (*genproto.CourierLocation, error) {
	return c.storage.CourierLocation().GetCourierLocation(req)
}

func (c *CourierLocationService) UpdateCourierLocation(ctx context.Context,  req *genproto.UpdateCourierLocationRequest) (*genproto.Empty, error) {
	return c.storage.CourierLocation().UpdateCourierLocation(req)
}

func (c *CourierLocationService) DeleteCourierLocation(ctx context.Context, req *genproto.DeleteCourierLocationRequest) (*genproto.Empty, error) {
	return c.storage.CourierLocation().DeleteCourierLocation(req)
}

func (c *CourierLocationService) ListCourierLocations(ctx context.Context, req *genproto.GetAllCourierLocationsRequest) (*genproto.CourierLocationList, error) {
	return c.storage.CourierLocation().ListCourierLocations(req)
}
