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
	es service.EventService
}

func NewTicketController(ts service.TicketService, es service.EventService) TicketController {
	return TicketController{
		ts: ts,
		es: es,
	}
}

func (tc TicketController) Create(c echo.Context) error {
	var ticket domain.Ticket
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	event, err := tc.checkEvent(c, uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
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

	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("ticket not found"))
	}

	event, err := tc.checkEvent(c, uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
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
	ticket_id, _ := strconv.Atoi(c.Param("ticket_id"))

	ticket, err := tc.ts.Get(uint(ticket_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if ticket.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("ticket not found"))
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

func (tc TicketController) checkEvent(c echo.Context, event_id uint) (domain.Event, error) {
	event, err := tc.es.Get(uint(event_id))
	if err != nil {
		return domain.Event{}, err
	}
	if event.ID == 0 {
		return domain.Event{}, errors.New("event not found")
	}
	return event, nil
}
