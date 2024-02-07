package main

import (
	"context"
	"fmt"

	"github.com/davidandw190/orders-microservice-api/service"
)

func main() {
	service := service.New()

	err := service.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start service:", err)
	}
}
