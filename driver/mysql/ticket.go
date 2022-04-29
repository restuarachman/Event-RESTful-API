package mysql

import (
	"ticketing/model/domain"

	"gorm.io/gorm"
)

type DBTicketService struct {
	db *gorm.DB
}

func NewDBTicketService(db *gorm.DB) *DBTicketService {
	return &DBTicketService{
		db: db,
	}
}

func (ts DBTicketService) Save(ticket domain.Ticket) (domain.Ticket, error) {
	tx := ts.db.Save(&ticket)
	err := tx.Error
	return ticket, err
}

func (ts DBTicketService) GetAll() ([]domain.Ticket, error) {
	tickets := []domain.Ticket{}
	tx := ts.db.Find(&tickets)
	err := tx.Error
	return tickets, err
}

func (ts DBTicketService) Get(id uint) (domain.Ticket, error) {
	ticket := domain.Ticket{}
	tx := ts.db.Find(&ticket, id)
	err := tx.Error
	return ticket, err
}

func (ts DBTicketService) Delete(ticket domain.Ticket) (domain.Ticket, error) {
	tx := ts.db.Delete(&ticket)
	err := tx.Error
	return ticket, err
}
func (ts DBTicketService) GetAllByEventId(event_id uint) ([]domain.Ticket, error) {
	tickets := []domain.Ticket{}
	tx := ts.db.Find(&tickets, "event_id = ?", event_id)
	err := tx.Error
	return tickets, err
}
