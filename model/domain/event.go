package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserId      uint      `json:"user id" form:"user id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	DateStart   time.Time `json:"date start" form:"date start"`
	DateEnd     time.Time `json:"date end" form:"date end"`
	Time        string    `json:"time" form:"time"`
	Location    string    `json:"location" form:"location"`
	User        User      `gorm:"foreignKey:UserId"`
	Tickets     []Ticket  `gorm:"foreignKey:EventId"`
}
