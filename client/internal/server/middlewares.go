package server

import (
	"context"
	"hetz-client/config"
	"hetz-client/internal/auth/token"
	"hetz-client/internal/repository"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// NewAuthMiddleware creates an Auth middleware with the given repository
// Yeah it's messy, but this has to be done when we don't have global repo variable :(
func NewAuthMiddleware(repo *repository.Repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := token.ValidateSession(c, repo); err != nil {
				return err
			}
			return next(c)
		}
	}
}

func LoadMiddlewares(e *echo.Echo, cfg *config.Config, repo *repository.Repository) {
	// Recover middleware
	// Prints panic stack trace and handles control to the centralized error handler
	e.Use(middleware.Recover())

	// Secure middleware
	// Protects against XSS, content type sniffing, clickjacking and insecure connections
	e.Use(middleware.Secure())

	// CORS middlware
	// Restricted to accept only requests from the same origin
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.BaseURL},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	// RequestLogger middleware
	// Uses log/slog and forwards error to global error handler to decide status code
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}

			return nil
		},
	}))

	// RateLimiter middleware
	// Limits the application to 10 requests/sec using the default in-memory store
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(10),
				Burst:     20,
				ExpiresIn: 2 * time.Minute,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			return ctx.RealIP(), nil
		},
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			return ctx.JSON(http.StatusTooManyRequests, nil)
		},
	}))

}
