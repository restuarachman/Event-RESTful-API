package service

import "ticketing/model/domain"

type EventService interface {
	Save(event domain.Event) (domain.Event, error)
	GetAll() ([]domain.Event, error)
	Get(id uint) (domain.Event, error)
	Delete(event domain.Event) (domain.Event, error)
}
