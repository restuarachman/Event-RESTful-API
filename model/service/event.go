package service

import "ticketing/model/domain"

type EventService interface {
	Add(event domain.Event, user_id uint) (domain.Event, error)
	GetAll() ([]domain.Event, error)
	Get(id uint) (domain.Event, error)
	Update(event_id uint, event domain.Event, jwtID uint) (domain.Event, error)
	Delete(event_id uint, jwtID uint) (domain.Event, error)
	GetByUserId(user_id uint) ([]domain.Event, error)
}
