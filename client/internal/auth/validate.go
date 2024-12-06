package auth

import (
	"fmt"
	"hetz-client/internal/models"
	"hetz-client/internal/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ValidateSession(c echo.Context, repo *repository.Repository) error {
	token, userId, err := getDataFromCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	fmt.Println("Token is: ", token)
	fmt.Println("User ID is: ", userId)
	session, err := repo.GetSessionToken(token)
	if err != nil {
		fmt.Println("Error getting session token: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if err := ValidateToken(session, token); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return nil
}

func getDataFromCookie(c echo.Context) (token string, userId int, err error) {
	cookieToken, err := c.Cookie("token")
	if err != nil || cookieToken == nil || cookieToken.Value == "" {
		return "", 0, err
	}

	userToken, err := c.Cookie("user_id")
	if err != nil {
		return "", 0, err
	}
	userId, err = strconv.Atoi(userToken.Value)
	if err != nil {
		return "", 0, err
	}

	return cookieToken.Value, userId, nil
}

func ValidateToken(session *models.SessionToken, cookieToken string) error {
	if session.ExpiresAt.Before(session.ExpiresAt) {
		fmt.Println("Session expired")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	if session.Token != cookieToken {
		fmt.Println("Token mismatch")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// take first 14 digits and convert them to time.Time
	created_at, err := time.Parse("20060102150405", session.Token[:14])
	if err != nil {
		fmt.Println("Error parsing created_at: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	fmt.Println("Created at: ", created_at)
	fmt.Println("Expires at: ", session.ExpiresAt)

	return nil
}
