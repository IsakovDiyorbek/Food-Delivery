package main

import (
	"fmt"
	"log"

	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/api"
	_ "github.com/Food_Delivery/Food-Delivery-Api-Gateway/docs"
	minio "github.com/Food_Delivery/Food-Delivery-Api-Gateway/minIO"
	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	delivery, err := grpc.NewClient(fmt.Sprintf("delivery_service%s", ":9000"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer delivery.Close()

	user, err := grpc.NewClient(fmt.Sprintf("auth_service%s", ":9999"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer user.Close()

	ConnMinIO, err := minio.InitMinioClient()
	if err != nil {
		log.Fatal(err)
	}

	enforcer, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewGin(user, delivery, ConnMinIO, enforcer)
	log.Println("starting server")
	r.Run(":8088")

}
