package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId       uint
	CheckoutDate time.Time
	Status       string `json:"status" form:"status"`
	TotalPrice   int
	User         User          `gorm:"foreignKey:UserId"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderId"`
}
