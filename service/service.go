package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *Service {
	return &Service{router: loadRoutes()}
}

func (s *Service) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: s.router,
	}

	err := s.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := s.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("starting server..")

	ch := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
