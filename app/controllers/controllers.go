package controllers

import (
	"net/http"

	"hetz/app/repository"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Up(ctx echo.Context) error {
	if err := c.repo.Up(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}
