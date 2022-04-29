package response

import "ticketing/model/domain"

type TicketResponse struct {
	EventId     uint
	Name        string
	Description string
	Price       int
}

func ToTicketResponse(ticket domain.Ticket) TicketResponse {
	return TicketResponse{
		EventId:     ticket.EventId,
		Name:        ticket.Name,
		Description: ticket.Description,
		Price:       ticket.Price,
	}
}

func ToTicketListResponse(tickets []domain.Ticket) []TicketResponse {
	response := []TicketResponse{}
	for _, value := range tickets {
		response = append(response, ToTicketResponse(value))
	}
	return response
}
