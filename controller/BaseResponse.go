package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Meta struct {
		Status  int
		Message string
	}
	Data interface{}
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := Response{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data
	return c.JSON(response.Meta.Status, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := Response{}
	response.Meta.Status = status
	response.Meta.Message = err.Error()
	return c.JSON(response.Meta.Status, response)
}
