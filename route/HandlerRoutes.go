package route

import (
	"net/http"
	"ticketing/controller"
	mid "ticketing/middleware"
	"ticketing/middleware/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteHandler struct {
	uc controller.UserController
	tc controller.TicketController
	oc controller.OrderController
	ec controller.EventController
}

func NewRouteHandler(
	uc controller.UserController,
	tc controller.TicketController,
	oc controller.OrderController,
	ec controller.EventController,
) RouteHandler {
	return RouteHandler{
		uc: uc,
		tc: tc,
		oc: oc,
		ec: ec,
	}
}

func (rh *RouteHandler) HandlerUserRoutes(e *echo.Echo) {
	// add log middleware
	mid.LogMiddleware(e)

	// add jwt middleware
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// add admin middleware
	eAdmin := eJwt.Group("")
	eAdmin.Use(mid.AdminMiddleware)

	// add customer middleware
	eCustomer := eJwt.Group("")
	eCustomer.Use(mid.CustomerMiddleware)

	// add event organizer middleware
	eEO := eJwt.Group("")
	eEO.Use(mid.EoMiddleware)

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	eAdmin.GET("api/v1/users", rh.uc.GetAll)
	eJwt.GET("api/v1/users/:user_id", rh.uc.Get)
	eJwt.PUT("api/v1/users/:user_id", rh.uc.Update)
	eJwt.DELETE("api/v1/users/:user_id", rh.uc.Delete)
	e.POST("api/v1/register", rh.uc.Register)
	e.POST("api/v1/login", rh.uc.Login)

	e.GET("api/v1/events", rh.ec.GetAll)
	eEO.POST("api/v1/events", rh.ec.Create)
	e.GET("api/v1/events/:event_id", rh.ec.Get)
	eEO.PUT("api/v1/events/:event_id", rh.ec.Update)
	eEO.DELETE("api/v1/events/:event_id", rh.ec.Delete)
	e.GET("api/v1/users/:user_id/events", rh.ec.GetAllEventByUserId)

	e.GET("api/v1/tickets", rh.tc.GetAll)
	e.GET("api/v1/tickets/:ticket_id", rh.tc.Get)
	e.GET("api/v1/events/:event_id/tickets", rh.tc.GetAllByEventId)
	eEO.POST("api/v1/events/:event_id/tickets", rh.tc.Create)
	eEO.PUT("api/v1/events/:event_id/tickets/:ticket_id", rh.tc.Update)
	eEO.DELETE("api/v1/events/:event_id/tickets/:ticket_id", rh.tc.Delete)

	eAdmin.GET("api/v1/orders", rh.oc.GetAll)
	eCustomer.POST("api/v1/orders", rh.oc.Create)
	eCustomer.GET("api/v1/orders/:order_id", rh.oc.Get)
	eCustomer.PUT("api/v1/orders/:order_id", rh.oc.Update)
	eCustomer.GET("api/v1/users/:user_id/orders", rh.oc.GetAllByUser)
}
