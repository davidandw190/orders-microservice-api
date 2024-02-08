package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/davidandw190/orders-microservice-api/service"
)

func main() {
	service := service.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := service.Start(ctx)
	if err != nil {
		fmt.Println("failed to start service:", err)
	}
}
