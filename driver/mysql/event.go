package mysql

import (
	"errors"
	"ticketing/model/domain"

	"gorm.io/gorm"
)

type DBEventService struct {
	db *gorm.DB
}

func NewDBEventService(db *gorm.DB) DBEventService {
	return DBEventService{
		db: db,
	}
}

func (es DBEventService) Add(event domain.Event, user_id uint) (domain.Event, error) {
	event.UserId = user_id
	tx := es.db.Save(&event)
	err := tx.Error
	return event, err
}

func (es DBEventService) GetAll() ([]domain.Event, error) {
	events := []domain.Event{}
	tx := es.db.Find(&events)
	err := tx.Error
	return events, err
}

func (es DBEventService) Get(id uint) (domain.Event, error) {
	event := domain.Event{}
	tx := es.db.Find(&event, id)
	err := tx.Error
	if event.ID == 0 {
		return event, errors.New("event not found")
	}
	return event, err
}

func (es DBEventService) Update(event_id uint, event domain.Event, jwtID uint) (domain.Event, error) {
	eventDB, err := es.Get(event_id)
	if err != nil {
		return domain.Event{}, err
	}
	if eventDB.UserId != jwtID {
		return domain.Event{}, errors.New("forbidden")
	}

	if event.Name != "" {
		eventDB.Name = event.Name
	}
	if event.Description != "" {
		eventDB.Description = event.Description
	}
	if !event.DateStart.IsZero() {
		eventDB.DateStart = event.DateStart
	}
	if !event.DateEnd.IsZero() {
		eventDB.DateEnd = event.DateEnd
	}
	if event.Time != "" {
		eventDB.Time = event.Time
	}
	if event.Location != "" {
		eventDB.Location = event.Location
	}
	tx := es.db.Save(&eventDB)
	err = tx.Error
	if err != nil {
		return domain.Event{}, err
	}
	return eventDB, err
}

func (es DBEventService) Delete(event_id uint, jwtID uint) (domain.Event, error) {
	eventDB, err := es.Get(event_id)
	if err != nil {
		return domain.Event{}, err
	}
	if eventDB.UserId != jwtID {
		return domain.Event{}, errors.New("forbidden")
	}
	tx := es.db.Delete(&eventDB)
	err = tx.Error
	return eventDB, err
}
func (es DBEventService) GetByUserId(user_id uint) ([]domain.Event, error) {
	events := []domain.Event{}
	tx := es.db.Find(&events, "user_id = ?", user_id)
	err := tx.Error
	return events, err
}
