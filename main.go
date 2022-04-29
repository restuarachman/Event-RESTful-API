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

	es := _serviceMYSQL.NewDBEventService(config.DB)
	ec := controller.NewEventController(es)

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
	eAdmin.GET("api/v1/users", uc.GetAll)
	eAdmin.GET("api/v1/users/:user_id", uc.Get)
	eCustomer.PUT("api/v1/users/:user_id", uc.Update)
	eCustomer.DELETE("api/v1/users/:user_id", uc.Delete)
	e.POST("api/v1/register", uc.Register)
	e.POST("api/v1/login", uc.Login)

	e.GET("api/v1/events", ec.GetAll)
	e.GET("api/v1/events/:event_id", ec.Get)
	eEO.GET("api/v1/users/:user_id/events", ec.GetAllEventByUserId, mid.SelfMiddleware)
	eEO.POST("api/v1/events", ec.Create1)
	eEO.POST("api/v1/users/:user_id/events", ec.Create2, mid.SelfMiddleware)
	eEO.GET("api/v1/users/:user_id/events/:event_id", ec.Get, mid.SelfMiddleware)
	// eEO.PUT("api/v1/events/:event_id", ec.Update)
	eEO.PUT("api/v1/users/:user_id/events/:event_id", ec.Update, mid.SelfMiddleware)
	// eEO.DELETE("api/v1/events/:event_id", ec.Delete)
	eEO.DELETE("api/v1/users/:user_id/events/:event_id", ec.Delete, mid.SelfMiddleware)

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
