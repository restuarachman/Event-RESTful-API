package main

import (
	"log"
	"net/http"
	"ticketing/config"
	"ticketing/controller"
	_serviceMYSQL "ticketing/driver/mysql"
	mid "ticketing/middleware"
	"ticketing/middleware/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	e := echo.New()

	// make controller
	us := _serviceMYSQL.NewDBUserService(config.DB)
	uc := controller.NewUserController(us)

	// add log middleware
	mid.LogMiddleware(e)
	// add jwt middlware
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eAdmin := eJwt.Group("")
	eAdmin.Use(mid.AdminMiddleware)

	eCustomer := eJwt.Group("")
	eCustomer.Use(mid.CustomerMiddleware)

	eEO := eJwt.Group("")
	eEO.Use(mid.EoMiddleware)

	// Route
	eAdmin.GET("api/v1/users", uc.GetAll, mid.AdminMiddleware)
	eAdmin.GET("api/v1/users/:id", uc.Get)
	eCustomer.PUT("api/v1/users/:id", uc.Edit)
	eCustomer.DELETE("api/v1/users/:id", uc.Delete)
	e.POST("api/v1/register", uc.Register)
	e.POST("api/v1/login", uc.Login)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
