package main

import (
	"log"
	"net/http"
	"ticketing/config"
	"ticketing/controller"
	"ticketing/route"

	_serviceMYSQL "ticketing/driver/mysql"

	"github.com/labstack/echo/v4"
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

	handlerRote := route.NewRouteHandler(uc, tc, oc, ec)
	handlerRote.HandlerUserRoutes(e)

	if err := e.Start(":8000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
