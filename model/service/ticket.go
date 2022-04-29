package service

import "ticketing/model/domain"

type TicketService interface {
	Save(ticket domain.Ticket) (domain.Ticket, error)
	GetAll() ([]domain.Ticket, error)
	Get(id uint) (domain.Ticket, error)
	Delete(ticket domain.Ticket) (domain.Ticket, error)
	GetByEventId(event_id uint) ([]domain.Ticket, error)
}
