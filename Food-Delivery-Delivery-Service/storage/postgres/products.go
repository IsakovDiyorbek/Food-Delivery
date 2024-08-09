package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food/Food-Delivery-Delivery-Service/helper"

	"github.com/google/uuid"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) CreateProduct(req *pb.CreateProductRequest) (*pb.ProductEmpty, error) {
	id := uuid.NewString()
	query := "INSERT INTO products(id, name, description, price, image_url) VALUES($1, $2, $3, $4, $5)"
	_, err := p.db.Exec(query, id, req.Name, req.Description, req.Price, req.ImageUrl)
	if err != nil {
		log.Fatal("Error while creating product", err)
	}
	return &pb.ProductEmpty{}, nil
}

func (p *ProductRepo) GetProduct(req *pb.GetProductRequest) (*pb.Product, error) {
	query := "SELECT * FROM products WHERE id = $1 and deleted_at = 0"
	var product pb.Product
	err := p.db.QueryRow(query, req.GetId()).Scan(
		&product.Id, &product.Name, &product.Description, &product.Price, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err != nil {
		log.Fatal("Error while get product", err)
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepo) UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductEmpty, error) {
	var args []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		args = append(args, req.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}
	if req.Description != "" && req.Description != "string" {
		args = append(args, req.Description)
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)))
	}
	if req.Price != 0 {
		args = append(args, req.Price)
		conditions = append(conditions, fmt.Sprintf("price = $%d", len(args)))
	}
	if req.ImageUrl != "" && req.ImageUrl != "string" {
		args = append(args, req.ImageUrl)
		conditions = append(conditions, fmt.Sprintf("image_url = $%d", len(args)))
	}
	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE products SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := p.db.Exec(query, args...)
	if err != nil {
		log.Fatal("Error while update product", err)
		return nil, err
	}

	return &pb.ProductEmpty{}, nil
}

func (p *ProductRepo) DeleteProduct(req *pb.DeleteProductRequest) (*pb.ProductEmpty, error) {
	query := "UPDATE products set deleted_at = $1 WHERE id = $2"
	_, err := p.db.Exec(query, time.Now().Unix(), req.GetId())
	if err != nil {
		log.Fatal("Error while delete product", err)
		return nil, err
	}
	return &pb.ProductEmpty{}, nil
}

func (p *ProductRepo) ListProducts(req *pb.GetAllProductsRequest) (*pb.ProductList, error) {
	query := `SELECT * FROM products `

	param := make(map[string]interface{})
	filter := `where deleted_at = 0`

	if len(req.Name) > 0 {
		param["name"] = req.Name
		filter += ` and name  =:name`
	}
	if len(req.Description) > 0 {
		param["description"] = req.Description
		filter += ` and description  =:description`
	}
	if len(req.ImageUrl) > 0 {
		param["image_url"] = req.ImageUrl
		filter += ` and image_url  =:image_url`
	}
	if req.Price != 0 {
		param["price"] = req.Price
		filter += ` and price  =:price`
	}

	query += filter

	query, arr := helper.ReplaceQueryParams(query, param)

	rows, err := p.db.Query(query, arr...)

	if err != nil {
		log.Fatal("Error while get products", err)
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Product
	for rows.Next() {
		var product pb.Product
		err := rows.Scan(
			&product.Id, &product.Name, &product.Description, &product.Price, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			log.Fatal("Error while scan products", err)
			return nil, err
		}
		products = append(products, &product)
	}
	return &pb.ProductList{Products: products}, nil
}
