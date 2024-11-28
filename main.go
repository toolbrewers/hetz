package main

import (
	"fmt"
	"net/http"
	"os"

	"hetz/app/controllers"
	"hetz/app/repository"
	"hetz/config"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := config.Load(); err != nil {
		panic(fmt.Errorf("failed to load app.yml or set env vars: %w", err))
	}

	repo, err := repository.New()
	if err != nil {
		panic(fmt.Errorf("failed to initialize repository: %w", err))
	}

	e := echo.New()
	e.HTTPErrorHandler = config.HTTPErrorHandler

	config.LoadMiddlewares(e)
	config.LoadRoutes(e, controllers.New(repo))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000" // Default port
	}
	if err := e.Start(port); err != http.ErrServerClosed {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
