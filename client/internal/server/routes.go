package server

import (
	"hetz-client/internal/controllers"
	"hetz-client/internal/repository"

	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, ctrl *controllers.Controller, repo *repository.Repository) {
	// Registers routes to serve static files such as images, javascript, html,
	// css, pdf, fonts and so on
	e.Static("/", "public")
	e.Static("/assets", "internal/assets")

	// Group for routes that require authentication
	authGroup := e.Group("")
	authGroup.Use(NewAuthMiddleware(repo))

	// Reveal health status on /up that returns 200 if the app boots without errors,
	// otherwise returns 500. Can be used by load balancers and uptime monitors to
	// verify that the app is live
	authGroup.GET("/up", ctrl.Up)

	// User authentication...
	e.POST("/signup", ctrl.Signup)
}
