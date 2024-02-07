package service

import (
	"context"
	"fmt"
	"net/http"
)

type Service struct {
	router http.Handler
}

func New() *Service {
	return &Service{router: loadRoutes()}
}

func (s *Service) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: s.router,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
