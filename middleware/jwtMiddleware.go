package middleware

import (
	"ticketing/middleware/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUser(c echo.Context) (uint, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return uint(userId), role
	}
	return 0, ""
}

// func ExtractTokenUserRole(c echo.Context) string {
// 	user := c.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		role := claims["role"].(string)

// 		return role
// 	}
// 	return ""
// }

// func ExtractTokenUserId(c echo.Context) uint {
// 	user := c.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		role := claims["userId"].(float64)

// 		return uint(role)
// 	}
// 	return 0
// }