package controllers

import (
	"net/http"

	"hetz-client/config"
	"hetz-client/internal/repository"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	cfg  *config.Config
	repo *repository.Repository
}

func New(cfg *config.Config, repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Up(ctx echo.Context) error {
	if err := c.repo.Up(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}
