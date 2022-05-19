package controller

import (
	"net/http"
	"strconv"
	mid "ticketing/middleware"
	"ticketing/model/domain"
	"ticketing/model/response"
	"ticketing/model/service"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	es service.EventService
}

func NewEventController(es service.EventService) EventController {
	return EventController{
		es: es,
	}
}

func (ec EventController) Create(c echo.Context) error {
	var event domain.Event

	c.Bind(&event)

	user_id, _ := mid.ExtractTokenUser(c)

	event, err := ec.es.Add(event, user_id)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToEventResponse(event))
}

func (ec EventController) GetAll(c echo.Context) error {
	events, err := ec.es.GetAll()
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToEventListResponse(events))
}

func (ec EventController) Get(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	event, err := ec.es.Get(uint(event_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToEventResponse(event))
}

func (ec EventController) Update(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))

	event := domain.Event{}
	c.Bind(&event)

	user_id, _ := mid.ExtractTokenUser(c)
	event, err := ec.es.Update(uint(event_id), event, user_id)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToEventResponse(event))
}

func (ec EventController) Delete(c echo.Context) error {
	event_id, _ := strconv.Atoi(c.Param("event_id"))
	user_id, _ := mid.ExtractTokenUser(c)

	event, err := ec.es.Delete(uint(event_id), user_id)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToEventResponse(event))
}

func (ec EventController) GetAllEventByUserId(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	events, err := ec.es.GetByUserId(uint(user_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToEventListResponse(events))
}

func (ec EventController) IsMyEvent(c echo.Context, event domain.Event) bool {
	user_id, _ := mid.ExtractTokenUser(c)

	return user_id == event.UserId
}
