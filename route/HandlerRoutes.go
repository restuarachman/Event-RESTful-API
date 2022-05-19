package route

import (
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
	jwtMiddlware := middleware.JWT([]byte(constants.SECRET_JWT))

	// Route
	e.GET("api/v1/users", rh.uc.GetAll, jwtMiddlware, mid.AdminMiddleware)
	e.GET("api/v1/users/:user_id", rh.uc.Get, jwtMiddlware)
	e.PUT("api/v1/users/:user_id", rh.uc.Update, jwtMiddlware)
	e.DELETE("api/v1/users/:user_id", rh.uc.Delete, jwtMiddlware)
	e.POST("api/v1/register", rh.uc.Register)
	e.POST("api/v1/login", rh.uc.Login)

	e.GET("api/v1/events", rh.ec.GetAll)
	e.POST("api/v1/events", rh.ec.Create, jwtMiddlware, mid.EoMiddleware)
	e.GET("api/v1/events/:event_id", rh.ec.Get)
	e.PUT("api/v1/events/:event_id", rh.ec.Update, jwtMiddlware, mid.EoMiddleware)
	e.DELETE("api/v1/events/:event_id", rh.ec.Delete, jwtMiddlware, mid.EoMiddleware)
	e.GET("api/v1/users/:user_id/events", rh.ec.GetAllEventByUserId)

	e.GET("api/v1/tickets", rh.tc.GetAll)
	e.GET("api/v1/tickets/:ticket_id", rh.tc.Get)
	e.GET("api/v1/events/:event_id/tickets", rh.tc.GetAllByEventId)
	e.POST("api/v1/events/:event_id/tickets", rh.tc.Create, jwtMiddlware, mid.EoMiddleware)
	e.PUT("api/v1/events/:event_id/tickets/:ticket_id", rh.tc.Update, jwtMiddlware, mid.EoMiddleware)
	e.DELETE("api/v1/events/:event_id/tickets/:ticket_id", rh.tc.Delete, jwtMiddlware, mid.EoMiddleware)

	e.GET("api/v1/orders", rh.oc.GetAll, jwtMiddlware, mid.AdminMiddleware)
	e.POST("api/v1/orders", rh.oc.Create, jwtMiddlware, mid.CustomerMiddleware)
	e.GET("api/v1/orders/:order_id", rh.oc.Get, jwtMiddlware, mid.CustomerMiddleware)
	e.PUT("api/v1/orders/:order_id", rh.oc.Update, jwtMiddlware, mid.CustomerMiddleware)
	e.GET("api/v1/users/:user_id/orders", rh.oc.GetAllByUser, jwtMiddlware, mid.CustomerMiddleware)
}
