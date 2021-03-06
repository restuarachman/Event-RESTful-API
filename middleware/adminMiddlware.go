package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := ExtractTokenUser(c)
		if role != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorize")
		}
		return next(c)
	}
}
