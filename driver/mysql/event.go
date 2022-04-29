package mysql

import (
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

func (es DBEventService) Save(event domain.Event) (domain.Event, error) {
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
	return event, err
}
func (es DBEventService) Delete(event domain.Event) (domain.Event, error) {
	tx := es.db.Delete(&event)
	err := tx.Error
	return event, err
}
func (es DBEventService) GetByUserId(user_id uint) ([]domain.Event, error) {
	events := []domain.Event{}
	tx := es.db.Find(&events, "user_id = ?", user_id)
	err := tx.Error
	return events, err
}
