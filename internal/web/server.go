package web

import (
	"context"
	"log"
	"net/http"
	"service/internal/config"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	ServerHTTP struct {
		address     string
		idleTimeout time.Duration
	}

	Server interface {
		Start(e *echo.Echo)
		Shutdown(e *echo.Echo)
	}
)

func NewServer(cfg config.Config) Server {
	return &ServerHTTP{
		address:     cfg.GetAddress(),
		idleTimeout: cfg.GetIdleTime(),
	}
}

func (s *ServerHTTP) Start(e *echo.Echo) {
	log.Printf("Server starting: %s...", s.address)
	if err := e.Start(s.address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func (s *ServerHTTP) Shutdown(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), s.idleTimeout)
	defer cancel()
	log.Printf("Server shutting down...")
	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}
	log.Println("Server by ended")
}
