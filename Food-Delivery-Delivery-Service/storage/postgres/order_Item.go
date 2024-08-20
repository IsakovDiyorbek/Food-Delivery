package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
	"github.com/google/uuid"
)

type OrderItemRepo struct {
	db *sql.DB
}

func NewOrderItemRepo(db *sql.DB) *OrderItemRepo {
	return &OrderItemRepo{db: db}
}

func (o *OrderItemRepo) CreateOrderItem(req *pb.CreateOrderItemRequest) (*pb.OrderItemEmpty, error) {
	query := `insert into order_items(id, order_id, product_id, quantity, price) values($1, $2, $3, $4, $5)`
	id := uuid.NewString()
	_, err := o.db.Exec(query, id, req.OrderId, req.ProductId, req.Quantity, req.Price)
	if err != nil {
		log.Fatal("Error while creating order item", err)
		return nil, err
	}
	return &pb.OrderItemEmpty{}, nil
}

func (o *OrderItemRepo) GetOrderItem(req *pb.GetOrderItemRequest) (*pb.OrderItem, error) {
	query := `
	SELECT
		oi.id, 
		oi.order_id, 
		oi.product_id, 
		oi.quantity, 
		oi.price AS order_item_price,
		oi.created_at AS order_item_created_at,
		p.id AS product_id,
		p.name AS product_name,
		p.description AS product_description,
		p.price AS product_price,
		p.image_url,
		p.created_at AS product_created_at,
		o.id AS order_id,
		o.user_id,
		o.courier_id,
		o.status,
		o.total_amount,
		o.delivery_address,
		o.created_at AS order_created_at
	FROM
		order_items oi
	JOIN 
		products p ON oi.product_id = p.id
	JOIN
		orders o ON oi.order_id = o.id
	WHERE
		oi.id = $1
	AND 
		o.deleted_at = 0;
`

	row := o.db.QueryRow(query, req.Id)
	

	var orderItem pb.OrderItem
	var order pb.Order
	var product pb.Product

	err := row.Scan(
		&orderItem.Id,
		&order.Id,
		&product.Id,
		&orderItem.Quantity,
		&orderItem.Price,
		&orderItem.CreatedAt,
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageUrl,
		&product.CreatedAt,
		&order.Id,
		&order.UserId,
		&order.CourierId,
		&order.Status,
		&order.TotalAmount,
		&order.DeliveryAddress,
		&order.CreatedAt,
	)

	if err != nil {
		log.Fatal("Error while scan order item", err)
		return nil, err
	}
	orderItem.OrderId = &order
	orderItem.ProductId = &product

	return &orderItem, nil

}

func (o *OrderItemRepo) UpdateOrderItem(req *pb.UpdateOrderItemRequest) (*pb.OrderItemEmpty, error) {
	var args []interface{}
	var conditions []string

	if req.OrderId != "" && req.OrderId != "string" {
		args = append(args, req.OrderId)
		conditions = append(conditions, fmt.Sprintf("order_id = $%d", len(args)))
	}
	if req.ProductId != "" && req.ProductId != "string" {
		args = append(args, req.ProductId)
		conditions = append(conditions, fmt.Sprintf("product_id = $%d", len(args)))
	}
	if req.Quantity != 0 {
		args = append(args, req.Quantity)
		conditions = append(conditions, fmt.Sprintf("quantity = $%d", len(args)))
	}
	if req.Price != 0 {
		args = append(args, req.Price)
		conditions = append(conditions, fmt.Sprintf("price = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE order_items SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := o.db.Exec(query, args...)
	if err != nil {
		log.Fatal("Error while updating order item", err)
		return nil, err
	}

	return &pb.OrderItemEmpty{}, nil

}

func (o *OrderItemRepo) ListOrderItems(req *pb.GetAllOrderItemsRequest) (*pb.OrderItemList, error) {
	query := `
		SELECT
			id, 
			order_id, 
			product_id, 
			quantity, 
			price,
			created_at
		FROM
			order_items
		`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.OrderId != "" {
		filters = append(filters, fmt.Sprintf("order_id = $%d", argCount))
		args = append(args, req.OrderId)
		argCount++
	}

	if req.ProductId != "" {
		filters = append(filters, fmt.Sprintf("product_id = $%d", argCount))
		args = append(args, req.ProductId)
		argCount++
	}

	if req.Quantity != 0 {
		filters = append(filters, fmt.Sprintf("quantity = $%d", argCount))
		args = append(args, req.Quantity)
		argCount++
	}

	if req.Price != 0 {
		filters = append(filters, fmt.Sprintf("price = $%d", argCount))
		args = append(args, req.Price)
		argCount++
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, req.Limit)
		argCount++

		if req.Ofset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Ofset)
			argCount++
		}
	}

	rows, err := o.db.Query(query, args...)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orderItems []*pb.OrderItem
	for rows.Next() {
		var orderItem pb.OrderItem
		var order pb.Order
		var product pb.Product
		err := rows.Scan(
			&orderItem.Id,
			&order.Id,
			&product.Id,
			&orderItem.Quantity,
			&orderItem.Price,
			&orderItem.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		orderItem.OrderId = &order
		orderItem.ProductId = &product
		orderItems = append(orderItems, &orderItem)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v\n", err)
		return nil, err
	}

	return &pb.OrderItemList{OrderItems: orderItems}, nil
}


func (o *OrderItemRepo) DeleteOrderItem(req *pb.DeleteOrderItemRequest) (*pb.OrderItemEmpty, error) {
	query := `
	UPDATE
		order_items
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := o.db.Exec(query, req.Id)

	if err != nil {
		log.Fatal("Error while deleting order item", err)
		return nil, err
	}
	return &pb.OrderItemEmpty{}, nil

}
