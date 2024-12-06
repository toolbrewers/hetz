package auth

import (
	"hetz-client/internal/models"
	"hetz-client/internal/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ValidateSession(c echo.Context, repo *repository.Repository) error {
	token, _, err := getDataFromCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	session, err := repo.GetSessionToken(token)
	if err != nil {
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

// TODO: Move most of this to the DB query
func ValidateToken(session *models.SessionToken, cookieToken string) error {
	if session.ExpiresAt.Before(session.ExpiresAt) {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	if session.Token != cookieToken {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// take first 14 digits and convert them to time.Time
	created_at, err := time.Parse("20060102150405", session.Token[:14])
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	// check if session.ExpiresAt is within 1 minute of first 14 digits of token
	bufferBefore := created_at.Add(-1 * time.Minute)
	bufferAfter := created_at.Add(1 * time.Minute)
	if session.ExpiresAt.Before(bufferBefore) || session.ExpiresAt.After(bufferAfter) {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return nil
}
