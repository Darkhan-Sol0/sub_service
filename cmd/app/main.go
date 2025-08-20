package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"service/internal/config"
	"service/internal/datasource/database"
	"service/internal/datasource/repository"
	"service/internal/service"
	"service/internal/web"
	"service/logger"

	"syscall"

	"github.com/labstack/echo/v4"

	_ "service/docs"
)

// @title Subscription Service API
// @version 1.0
// @description API for managing subscriptions
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	cfg := config.LoadConfig()
	s := web.NewServer(cfg)

	db, err := database.ConnectDB(context.Background(), cfg)
	if err != nil {
		log.Fatalln("error connect db: %w", err)
	}

	r := web.NewRouting(service.NewService(repository.NewDatabase(db)), logger.Init(cfg))
	r.RegisterRoutes(e)

	go func() {
		s.Start(e)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	s.Shutdown(e)
}
