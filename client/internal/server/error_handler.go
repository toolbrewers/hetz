package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, ctx echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		switch he.Code {
		case
			http.StatusNotFound,
			http.StatusNotAcceptable,
			http.StatusUnprocessableEntity,
			http.StatusInternalServerError:
			if err := ctx.Redirect(
				http.StatusSeeOther, fmt.Sprintf("/%d.html", he.Code)); err != nil {
				ctx.Logger().Error(err)
			}
		default:
			ctx.NoContent(he.Code)
		}
	} else {
		ctx.NoContent(http.StatusInternalServerError)
	}
}
