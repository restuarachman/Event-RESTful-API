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
	DB := config.ConnectDB()
	e := echo.New()

	// make controller
	us := _serviceMYSQL.NewDBUserService(DB)
	uc := controller.NewUserController(us)

	es := _serviceMYSQL.NewDBEventService(DB)
	ec := controller.NewEventController(es)

	ts := _serviceMYSQL.NewDBTicketService(DB)
	tc := controller.NewTicketController(ts, ec)

	os := _serviceMYSQL.NewDBOrderService(DB)
	ods := _serviceMYSQL.NewDBOrderDetailService(DB)
	oc := controller.NewOrderController(os, ods, tc)

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
	eJwt.PUT("api/v1/users/:user_id", uc.Update)
	eJwt.DELETE("api/v1/users/:user_id", uc.Delete)
	e.POST("api/v1/register", uc.Register)
	e.POST("api/v1/login", uc.Login)

	e.GET("api/v1/events", ec.GetAll)
	eEO.POST("api/v1/events", ec.Create)
	e.GET("api/v1/events/:event_id", ec.Get)
	eEO.PUT("api/v1/events/:event_id", ec.Update)
	eEO.DELETE("api/v1/events/:event_id", ec.Delete)
	e.GET("api/v1/users/:user_id/events", ec.GetAllEventByUserId)

	e.GET("api/v1/tickets", tc.GetAll)
	e.GET("api/v1/tickets/:ticket_id", tc.Get)
	e.GET("api/v1/events/:event_id/tickets", tc.GetAllByEventId)
	eEO.POST("api/v1/events/:event_id/tickets", tc.Create)
	eEO.PUT("api/v1/events/:event_id/tickets/:ticket_id", tc.Update)
	eEO.DELETE("api/v1/events/:event_id/tickets/:ticket_id", tc.Delete)

	eAdmin.GET("api/v1/orders", oc.GetAll)
	eCustomer.POST("api/v1/orders", oc.Create)
	eCustomer.GET("api/v1/orders/:order_id", oc.Get)
	eCustomer.PUT("api/v1/orders/:order_id", oc.Update)
	eCustomer.GET("api/v1/users/:user_id/orders", oc.GetAllByUser)

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
