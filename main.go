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

	repo := repository.New(repository.NewConfig())
	if err := repo.Up(); err != nil {
		panic(fmt.Errorf("failed to initialize repository: %w", err))
	}

	e := echo.New()
	e.HTTPErrorHandler = config.HTTPErrorHandler

	config.LoadMiddlewares(e)
	config.LoadRoutes(e, controllers.New(repo))

	if err := e.Start(os.Getenv("PORT")); err != http.ErrServerClosed {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
