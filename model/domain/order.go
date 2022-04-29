package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId       int
	CheckoutDate string
	Status       string
	TotalPrice   int
	User         User           `gorm:"foreignKey:UserId"`
	OrderDetails []OrderDetails `gorm:"foreignKey:OrderId"`
}
