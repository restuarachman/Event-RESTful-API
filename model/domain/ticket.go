package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventId      uint           `json:"event id" form:"event id"`
	Name         string         `json:"name" form:"name"`
	Description  string         `json:"description" form:"description"`
	Price        int            `json:"price" form:"price"`
	Event        Event          `gorm:"foreignKey:EventId"`
	OrderDetails []OrderDetails `gorm:"foreignKey:TicketId"`
}
