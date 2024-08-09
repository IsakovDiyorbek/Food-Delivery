package storage

import (
	pb "github.com/Food/Food-Delivery-Delivery-Service/genproto"
)

type StorageI interface {
	Product() ProductI
	Order() OrderI
	Cart() CartI
	OrderItem() OrderItemI
	Task() TaskI
	CourierLocation() CourierLocationI
	Notification() NotificationI
}

type ProductI interface {
	CreateProduct(req *pb.CreateProductRequest) (*pb.ProductEmpty, error)
	GetProduct(req *pb.GetProductRequest) (*pb.Product, error)
	UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductEmpty, error)
	DeleteProduct(req *pb.DeleteProductRequest) (*pb.ProductEmpty, error)
	ListProducts(req *pb.GetAllProductsRequest) (*pb.ProductList, error)
}

type OrderI interface {
	CreateOrder(req *pb.CreateOrderRequest) (*pb.OrderEmpty, error)
	GetOrder(req *pb.GetOrderRequest) (*pb.Order, error)
	UpdateOrder(req *pb.UpdateOrderRequest) (*pb.OrderEmpty, error)
	ListOrders(req *pb.GetAllOrdersRequest) (*pb.OrderList, error)
	DeleteOrder(req *pb.DeleteOrderRequest) (*pb.OrderEmpty, error)
}

type CartI interface {
	CreateCart(req *pb.CreateCartReq) (*pb.CartEmpty, error)
	GetCart(req *pb.GetByIdCartRequest) (*pb.Cart, error)
	GetAllCarts(req *pb.GetAllCartsReq) (*pb.GetAllCartsRes, error)
	UpdateCart(req *pb.UpdateCartReq) (*pb.UpdateCartRes, error)
	DeleteCart(req *pb.DeleteCartRequest) (*pb.DeleteCartResp, error)
}

type OrderItemI interface {
	CreateOrderItem(req *pb.CreateOrderItemRequest) (*pb.OrderItemEmpty, error)
	GetOrderItem(req *pb.GetOrderItemRequest) (*pb.OrderItem, error)
	UpdateOrderItem(req *pb.UpdateOrderItemRequest) (*pb.OrderItemEmpty, error)
	DeleteOrderItem(req *pb.DeleteOrderItemRequest) (*pb.OrderItemEmpty, error)
	ListOrderItems(req *pb.GetAllOrderItemsRequest) (*pb.OrderItemList, error)
}

type TaskI interface {
	CreateTask(req *pb.CreateTaskRequest) (*pb.TaskEmpty, error)
	GetTask(req *pb.GetTaskRequest) (*pb.Task, error)
	UpdateTask(req *pb.UpdateTaskRequest) (*pb.TaskEmpty, error)
	DeleteTask(req *pb.DeleteTaskRequest) (*pb.TaskEmpty, error)
	ListTasks(req *pb.GetAllTasksRequest) (*pb.TaskList, error)
}

type CourierLocationI interface {
	CreateCourierLocation(req *pb.CreateCourierLocationRequest) (*pb.Empty, error)
	GetCourierLocation(req *pb.GetCourierLocationRequest) (*pb.CourierLocation, error)
	UpdateCourierLocation(req *pb.UpdateCourierLocationRequest) (*pb.Empty, error)
	DeleteCourierLocation(req *pb.DeleteCourierLocationRequest) (*pb.Empty, error)
	ListCourierLocations(req *pb.GetAllCourierLocationsRequest) (*pb.CourierLocationList, error)
}


type NotificationI interface{
	CreateNotification(req *pb.CreateNotificationRequest) (*pb.NotificationEmpty, error)
	GetNotification(req *pb.GetNotificationRequest) (*pb.Notification, error)
	MarkNotificationAsRead(req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error)
}