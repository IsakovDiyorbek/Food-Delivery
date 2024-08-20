package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Food-Delivery/Food-Delivery-Delivery-Service/helper"

	pb "github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"

	"github.com/google/uuid"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (o *OrderRepo) CreateOrder(req *pb.CreateOrderRequest) (*pb.OrderEmpty, error) {
	query := "INSERT INTO orders(id, user_id, courier_id, status, total_amount, delivery_address) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := o.db.Exec(query, uuid.NewString(), req.GetUserId(), req.GetCourierId(), req.GetStatus(), req.GetTotalAmount(), req.GetDeliveryAddress())
	if err != nil {
		log.Printf("Error while creating order: %v", err)
		return nil, err
	}
	return &pb.OrderEmpty{}, nil
}

func (o *OrderRepo) GetOrder(req *pb.GetOrderRequest) (*pb.Order, error) {
	query := "SELECT id, user_id, courier_id, status, total_amount, delivery_address, created_at FROM orders WHERE id = $1 and deleted_at = 0"
	var order pb.Order
	err := o.db.QueryRow(query, req.GetId()).Scan(
		&order.Id, &order.UserId, &order.CourierId, &order.Status, &order.TotalAmount, &order.DeliveryAddress, &order.CreatedAt)
	if err != nil {
		log.Fatal("Error while get order", err)
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) UpdateOrder(req *pb.UpdateOrderRequest) (*pb.OrderEmpty, error) {
	var args []interface{}
	var conditions []string

	if req.UserId != "" && req.UserId != "string" {
		args = append(args, req.UserId)
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)))
	}
	if req.CourierId != "" && req.CourierId != "string" {
		args = append(args, req.CourierId)
		conditions = append(conditions, fmt.Sprintf("courier_id = $%d", len(args)))
	}
	if req.TotalAmount != 0 {
		args = append(args, req.TotalAmount)
		conditions = append(conditions, fmt.Sprintf("total_amount = $%d", len(args)))
	}
	if req.Status != "" && req.Status != "string" {
		args = append(args, req.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}
	if req.DeliveryAddress != "" && req.DeliveryAddress != "string" {
		args = append(args, req.DeliveryAddress)
		conditions = append(conditions, fmt.Sprintf("delivery_address = $%d", len(args)))
	}
	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE orders SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	tx, err := o.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.OrderEmpty{}, nil
}

func (o *OrderRepo) DeleteOrder(req *pb.DeleteOrderRequest) (*pb.OrderEmpty, error) {
	query := "update orders set deleted_at = $1 WHERE id = $2"
	_, err := o.db.Exec(query, time.Now().Unix(), req.GetId())
	if err != nil {
		log.Fatal("Error while delete order", err)
		return nil, err
	}
	return &pb.OrderEmpty{}, nil
}
func (o *OrderRepo) ListOrders(req *pb.GetAllOrdersRequest) (*pb.OrderList, error) {
	query := `SELECT id, user_id, courier_id, status, total_amount, delivery_address, created_at FROM orders WHERE deleted_at = 0`

	param := make(map[string]interface{})

	if len(req.UserId) > 0 {
		param["user_id"] = req.UserId
		query += ` AND user_id = :user_id`
	}
	if len(req.CourierId) > 0 {
		param["courier_id"] = req.CourierId
		query += ` AND courier_id = :courier_id`
	}
	if len(req.Status) > 0 {
		param["status"] = req.Status
		query += ` AND status = :status`
	}
	if req.TotalAmount != 0 {
		param["total_amount"] = req.TotalAmount
		query += ` AND total_amount = :total_amount`
	}
	if len(req.DeliveryAddress) > 0 {
		param["delivery_address"] = req.DeliveryAddress
		query += ` AND delivery_address = :delivery_address`
	}

	query, arr := helper.ReplaceQueryParams(query, param)

	rows, err := o.db.Query(query, arr...)

	if err != nil {
		log.Fatal("Error while get orders", err)
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var order pb.Order
		err := rows.Scan(
			&order.Id, &order.UserId, &order.CourierId, &order.Status, &order.TotalAmount, &order.DeliveryAddress, &order.CreatedAt)
		if err != nil {
			log.Fatal("Error while scan orders", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &pb.OrderList{Orders: orders}, nil
}
