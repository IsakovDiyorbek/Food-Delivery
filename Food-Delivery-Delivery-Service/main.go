package main

import (
	"log"
	"net"
	"net/http"

	"github.com/Food/Food-Delivery-Delivery-Service/config"
	"github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food/Food-Delivery-Delivery-Service/service"
	"github.com/Food/Food-Delivery-Delivery-Service/storage/postgres"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	cfg := config.Load()

	db, err := postgres.DbConnection()
	if err != nil {
		log.Fatal("Error while db connection", err)
	}

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal("Error while listening", err)
	}

	s := grpc.NewServer()

	genproto.RegisterProductServiceServer(s, service.NewProductService(db))
	genproto.RegisterOrderServiceServer(s, service.NewOrderService(db))
	genproto.RegisterCartServiceServer(s, service.NewCartService(db))
	genproto.RegisterOrderItemServiceServer(s, service.NewOrderItemsService(db))
	genproto.RegisterTaskServiceServer(s, service.NewTaskService(db))
	genproto.RegisterCourierLocationServiceServer(s, service.NewCourierLocationService(db))
	genproto.RegisterNotificationServiceServer(s, service.NewNotificationService(db))


	http.HandleFunc("/ws", websocketHandler)
	go func() {
		log.Printf("HTTP server started on port %s", ":8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("gRPC server started on port %s", cfg.HTTPPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
