package server

import (
	"hetz-client/internal/controllers"

	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, ctrl *controllers.Controller) {
	// Registers routes to serve static files such as images, javascript, html,
	// css, pdf, fonts and so on
	e.Static("/", "public")
	e.Static("/assets", "internal/assets")

	// Reveal health status on /up that returns 200 if the app boots without errors,
	// otherwise returns 500. Can be used by load balancers and uptime monitors to
	// verify that the app is live
	e.GET("/up", ctrl.Up)

	// User authentication...
	e.POST("/signup", ctrl.Signup)
}
