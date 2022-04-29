package response

import (
	"ticketing/model/domain"
	"time"
)

type EventResponse struct {
	ID          uint
	UserId      uint
	Name        string
	Description string
	DateStart   time.Time
	DateEnd     time.Time
	Time        string
	Location    string
}

func ToEventResponse(event domain.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		UserId:      event.UserId,
		Name:        event.Name,
		Description: event.Description,
		DateStart:   event.DateStart,
		DateEnd:     event.DateEnd,
		Time:        event.Time,
		Location:    event.Location,
	}
}

func ToEventListResponse(events []domain.Event) []EventResponse {
	response := []EventResponse{}
	for _, value := range events {
		response = append(response, ToEventResponse(value))
	}
	return response
}
