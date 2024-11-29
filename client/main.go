package main

import (
	"fmt"

	"hetz-client/config"
	"hetz-client/internal/controllers"
	"hetz-client/internal/repository"
	"hetz-client/internal/server"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf("failed to load env file: %w", err))
	}

	cfg := config.New()
	repo := repository.New("db/app.db")
	ctrl := controllers.New(cfg, repo)

	e := echo.New()
	e.HTTPErrorHandler = server.HTTPErrorHandler

	server.LoadMiddlewares(e, cfg)
	server.LoadRoutes(e, ctrl)

	if err := e.Start(cfg.Port); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
