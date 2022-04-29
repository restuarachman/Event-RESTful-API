package middleware

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SelfMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("user_id"))
		tokernUserId, _ := ExtractTokenUser(c)
		if tokernUserId != uint(userId) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorize")
		}
		return next(c)
	}
}
