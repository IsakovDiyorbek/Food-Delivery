package service

import (
	"context"

	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/storage"
)

type NotificationService struct {
	storage storage.StorageI
	genproto.UnimplementedNotificationServiceServer
}

func NewNotificationService(storage storage.StorageI) *NotificationService {
	return &NotificationService{storage: storage}
}

func (n *NotificationService) CreateNotification(ctx context.Context, req *genproto.CreateNotificationRequest) (*genproto.NotificationEmpty, error) {
	return n.storage.Notification().CreateNotification(req)
}

func (n *NotificationService) GetNotification(ctx context.Context, req *genproto.GetNotificationRequest) (*genproto.Notification, error) {
	return n.storage.Notification().GetNotification(req)
}

func (n *NotificationService) MarkNotificationAsRead(ctx context.Context, req *genproto.MarkNotificationAsReadReq) (*genproto.MarkNotificationAsReadResp, error) {
	return n.storage.Notification().MarkNotificationAsRead(req)
}
