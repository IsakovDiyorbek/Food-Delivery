package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Food/Food-Delivery-Delivery-Service/config"
	"github.com/Food/Food-Delivery-Delivery-Service/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	db       *sql.DB
	Products storage.ProductI
	Orders   storage.OrderI
	Carts  storage.CartI
	OrderItems storage.OrderItemI
	Tasks storage.TaskI
	CourierLocations storage.CourierLocationI
	Notifications storage.NotificationI
}

func DbConnection() (storage.StorageI, error) {
	cfg := config.Load()
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal("Error while db connection", err)
		return nil, nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error while db ping connection", err)
		return nil, nil

	}
	return &Storage{db: db}, nil

}

func (s *Storage) Product() storage.ProductI {
	if s.Products == nil {
		s.Products = &ProductRepo{s.db}
	}
	return s.Products
}

func (s *Storage) Order() storage.OrderI {
	if s.Orders == nil {
		s.Orders = &OrderRepo{s.db}
	}
	return s.Orders
}


func (s *Storage) Cart() storage.CartI {
	if s.Carts == nil {
		s.Carts = &CartRepo{s.db}
	}
	return s.Carts
}


func (s *Storage) OrderItem() storage.OrderItemI {
	if s.OrderItems == nil {
		s.OrderItems = &OrderItemRepo{s.db}
	}
	return s.OrderItems
}

func (s *Storage) Task() storage.TaskI {
	if s.Tasks == nil {
		s.Tasks = &TasksRepo{s.db}
	}
	return s.Tasks
}

func (s *Storage) CourierLocation() storage.CourierLocationI {
	if s.CourierLocations == nil {
		s.CourierLocations = &CourierLocationRepo{s.db}
	}
	return s.CourierLocations
}

func (s *Storage) Notification() storage.NotificationI {
	if s.Notifications == nil {
		s.Notifications = &NotificationManager{s.db}
	}
	return s.Notifications
}