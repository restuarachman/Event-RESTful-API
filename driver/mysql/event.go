package mysql

import (
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

// func (es DBEventService) Save(event domain.Event) (domain.Event, error)
// func (es DBEventService) GetAll() ([]domain.Event, error)
// func (es DBEventService) Get(id uint) (domain.Event, error)
// func (es DBEventService) Delete(event domain.Event) (domain.Event, error)
