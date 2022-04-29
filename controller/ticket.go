package controller

import (
	"strconv"
	"ticketing/model/domain"
	"ticketing/model/service"

	"github.com/labstack/echo/v4"
)

type TicketController struct {
	ts service.TicketService
}

func NewTicketController(ts service.TicketService) TicketController {
	return TicketController{
		ts: ts,
	}
}

func (tc TicketController) Create1(c echo.Context) error {
	var ticket domain.Ticket
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	c.Bind(&ticket)
	ticket.EventId = uint(event_id)
	return nil
}
