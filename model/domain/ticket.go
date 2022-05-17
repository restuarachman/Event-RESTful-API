package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventId      uint
	Name         string        `json:"name" form:"name"`
	Description  string        `json:"description" form:"description"`
	Price        int           `json:"price" form:"price"`
	Event        Event         `gorm:"foreignKey:EventId"`
	OrderDetails []OrderDetail `gorm:"foreignKey:TicketId"`
}
