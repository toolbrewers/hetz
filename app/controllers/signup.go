package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"hetz/app/models"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Signup(ctx echo.Context) error {
	req := new(models.Signup)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	form, err := ctx.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	validate := validator.New()
	for key, _ := range form {
		if fn, ok := models.SignupValidations[key]; ok {
			if err := fn(validate, ctx.FormValue(key)); err != nil {
				return ctx.String(http.StatusBadRequest, models.SignupHelpers[key])
			}
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MaxCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	id, err := c.repo.CreateUser(&models.CreateUser{
		Username:     req.Username,
		Email:        req.Email,
		Password:     string(bytes),
		HetznerToken: req.HetznerToken,
	})

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": os.Getenv("APP_NAME"),
			"sub": strconv.Itoa(int(id)),
			"aud": []string{os.Getenv("BASE_URL")},
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"nbf": time.Now().Unix(),
			"iat": time.Now().Unix(),
		}).SignedString([]byte("supersecretkey")) // TODO: Use a cookie to set this value
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // TODO: Use a cookie to set this value
		HttpOnly: true,
		Path:     "/",
	})

	ctx.Response().Header().Add("HX-Redirect", "/dashboard")
	return ctx.NoContent(http.StatusCreated)
}
