package auth

import (
	"fmt"
	"hetz-client/internal/models"
	"hetz-client/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ValidateSession(c echo.Context, repo *repository.Repository) error {
	token, _, err := getDataFromCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	fmt.Println("Token is: ", token)

	// session, err := repo.GetSessionToken(token)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusUnauthorized)
	// }

	// if err := ValidateToken(session); err != nil {
	// 	return echo.NewHTTPError(http.StatusUnauthorized)
	// }

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

func ValidateToken(session *models.SessionToken) error {
	fmt.Println(session.DeletedAt)

	return nil
}
