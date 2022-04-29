package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing/model/domain"
	"ticketing/model/response"
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

func (tc TicketController) Create(c echo.Context) error {
	var ticket domain.Ticket
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	c.Bind(&ticket)
	ticket.EventId = uint(event_id)

	ticket, err := tc.ts.Save(ticket)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) GetAll(c echo.Context) error {
	tickets, err := tc.ts.GetAll()
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToTicketListResponse(tickets))
}

func (tc TicketController) Get(c echo.Context) error {
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))
	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) Update(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))

	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.Bind(&ticket)
	ticket.EventId = uint(event_id)

	ticket, err = tc.ts.Save(ticket)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) Delete(c echo.Context) error {
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))

	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("ticket not found"))
	}

	ticket, err = tc.ts.Delete(ticket)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) GetAllByEventId(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	tickets, err := tc.ts.GetAllByEventId(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToTicketListResponse(tickets))
}
