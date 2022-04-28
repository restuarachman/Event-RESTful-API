package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func EoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ExtractTokenUserRole(c)
		if role == "eo" || role == "admin" {
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorize")

	}
}
