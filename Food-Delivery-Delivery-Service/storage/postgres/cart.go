package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/google/uuid"
)

type CartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (c *CartRepo) CreateCart(req *pb.CreateCartReq) (*pb.CartEmpty, error) {
	option, err := json.Marshal(req.Options)
	if err != nil {
		log.Fatal("Error while marshal options", err)
	}

	query := `insert into cart(id, user_id, product_id, quantity, options) values($1, $2, $3, $4, $5)`
	_, err = c.db.Exec(query, uuid.NewString(), req.GetUserId(), req.Product, req.GetQuantity(), option)
	if err != nil {
		log.Fatal("Error while creating cart", err)
	}

	return &pb.CartEmpty{}, nil
}

func (c *CartRepo) GetCart(req *pb.GetByIdCartRequest) (*pb.Cart, error) {
	query := `SELECT
			c.id,
			c.user_id,
			c.product_id,
			c.quantity,
			c.options,
			c.created_at,
			p.name,
			p.description,
			p.price,
			p.image_url
		FROM
			cart c
		JOIN
			products p
		ON
			c.product_id = p.id
		WHERE
			c.id = $1
		AND 
			c.deleted_at = 0`
	row := c.db.QueryRow(query, req.GetId())
	var cart pb.Cart
	var product pb.Product
	err := row.Scan(
		&cart.Id,
		&cart.UserId,
		&product.Id,
		&cart.Quantity,
		&cart.Options,
		&cart.CreatedAt,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageUrl,
	)

	if err != nil {
		log.Fatal("Error while get cart", err)
		return nil, err
	}
	cart.Product = &product
	return &cart, nil
}

func (c *CartRepo) GetAllCarts(req *pb.GetAllCartsReq) (*pb.GetAllCartsRes, error) {
	query := `SELECT
            c.id,
            c.user_id,
            c.quantity,
            c.options,
            c.created_at,
			p.id,
            p.name,
            p.description,
            p.price,
            p.image_url
        FROM
            cart c
        JOIN
            products p 
        ON 
            c.product_id = p.id`
	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.Quantity != 0 {
		filters = append(filters, fmt.Sprintf("c.quantity = $%d", argCount))
		args = append(args, req.Quantity)
		argCount++
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, req.Limit)
		argCount++
	}
	if req.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, req.Offset)
		argCount++
	}

	rows, err := c.db.Query(query, args...)
	if err != nil {
		log.Println("Error while getting all carts", err)
		return nil, err
	}
	defer rows.Close()

	var carts []*pb.Cart
	for rows.Next() {
		var cart pb.Cart
		var product pb.Product
		err := rows.Scan(
			&cart.Id,
			&cart.UserId,
			&cart.Quantity,
			&cart.Options,
			&cart.CreatedAt,
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageUrl,
		)
		if err != nil {
			log.Fatal("Error while scan carts", err)
			return nil, err
		}
		cart.Product = &product
		carts = append(carts, &cart)
	}
	return &pb.GetAllCartsRes{Carts: carts}, nil
}

func (c *CartRepo) UpdateCart(req *pb.UpdateCartReq) (*pb.UpdateCartRes, error) {
	var args []interface{}
	var conditions []string

	optionsJSON, err := json.Marshal(req.Options)
	if err != nil {
		log.Println("Error while marshaling options to JSON", err)
		return nil, err
	}

	if req.UserId != "" && req.UserId != "string" {
		args = append(args, req.UserId)
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)))
	}
	if req.Product != "" && req.Product != "string" {
		args = append(args, req.Product)
		conditions = append(conditions, fmt.Sprintf("product_id = $%d", len(args)))
	}
	if req.Quantity != 0 {
		args = append(args, req.Quantity)
		conditions = append(conditions, fmt.Sprintf("quantity = $%d", len(args)))
	}
	if string(optionsJSON) != "" && string(optionsJSON) != "string" {
		args = append(args, string(optionsJSON))
		conditions = append(conditions, fmt.Sprintf("options = $%d", len(args)))
	}
	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE cart SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	tx, err := c.db.Begin()
	if err != nil {
		log.Fatal("Error while starting transaction", err)
		return nil, err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		log.Fatal("Error while updating cart", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("Error while commiting cart update", err)
		return nil, err
	}

	return &pb.UpdateCartRes{}, nil
}

func (c *CartRepo) DeleteCart(req *pb.DeleteCartRequest) (*pb.DeleteCartResp, error) {
	query := `
	UPDATE
		cart
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := c.db.Exec(query, req.Id)

	if err != nil {
		log.Fatal("Error while deleting cart", err)
		return nil, err
	}
	return &pb.DeleteCartResp{}, nil

}
