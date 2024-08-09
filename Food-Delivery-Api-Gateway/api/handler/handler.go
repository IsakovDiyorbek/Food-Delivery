package handler

import (
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"github.com/casbin/casbin/v2"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc"
)

type Handler struct {
	Product         genproto.ProductServiceClient
	User            user.UserServiceClient
	Order           genproto.OrderServiceClient
	OrderItems      genproto.OrderItemServiceClient
	Cart            genproto.CartServiceClient
	Task            genproto.TaskServiceClient
	CourierLocation genproto.CourierLocationServiceClient
	miniIO          *minio.Client
	Notification    genproto.NotificationServiceClient
	Enforcer        *casbin.Enforcer
}

func NewHandler(users, delivery *grpc.ClientConn, minIO *minio.Client, enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		Product:         genproto.NewProductServiceClient(delivery),
		User:            user.NewUserServiceClient(users),
		Order:           genproto.NewOrderServiceClient(delivery),
		OrderItems:      genproto.NewOrderItemServiceClient(delivery),
		Cart:            genproto.NewCartServiceClient(delivery),
		Task:            genproto.NewTaskServiceClient(delivery),
		CourierLocation: genproto.NewCourierLocationServiceClient(delivery),
		miniIO:          minIO,
		Notification:    genproto.NewNotificationServiceClient(delivery),
		Enforcer:        enforcer,
	}
}
