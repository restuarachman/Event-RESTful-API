package controller

import (
	"errors"
	"net/http"
	"strconv"
	"ticketing/model/domain"
	"ticketing/model/response"
	"ticketing/model/service"

	"github.com/labstack/echo/v4"
)

type TicketController struct {
	ts service.TicketService
	ec EventController
}

func NewTicketController(ts service.TicketService, ec EventController) TicketController {
	return TicketController{
		ts: ts,
		ec: ec,
	}
}

func (tc TicketController) Create(c echo.Context) error {
	var ticket domain.Ticket
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	// check Event
	event, err := tc.ec.es.Get(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if event.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, errors.New("Event not found"))
	}

	// check event is mine
	if !tc.ec.IsMyEvent(c, event) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	c.Bind(&ticket)
	ticket.EventId = uint(event.ID)

	ticket, err = tc.ts.Save(ticket)
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

	// check ticket
	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("ticket not found"))
	}

	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) Update(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))

	// check Event
	event, err := tc.ec.es.Get(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if event.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, errors.New("Event not found"))
	}

	// check Ticket
	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("ticket not found"))
	}

	// check Ticket && Event is mine
	if !tc.IsMyTicket(c, event, ticket) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	c.Bind(&ticket)
	ticket.EventId = uint(event.ID)

	ticket, err = tc.ts.Save(ticket)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) Delete(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))

	// check Event
	event, err := tc.ec.es.Get(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if event.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, errors.New("Event not found"))
	}

	// check Ticket
	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("ticket not found"))
	}

	// check Ticket && Event is mine
	if !tc.IsMyTicket(c, event, ticket) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	ticket, err = tc.ts.Delete(ticket)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToTicketResponse(ticket))
}

func (tc TicketController) GetAllByEventId(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	// check Event
	event, err := tc.ec.es.Get(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if event.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, errors.New("Event not found"))
	}

	tickets, err := tc.ts.GetAllByEventId(uint(event.ID))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToTicketListResponse(tickets))
}

func (tc TicketController) IsMyTicket(c echo.Context, event domain.Event, ticket domain.Ticket) bool {
	if tc.ec.IsMyEvent(c, event) && event.ID == ticket.EventId {
		return true
	}

	return false
}
