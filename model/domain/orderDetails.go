package domain

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	order_id int
	tiket_id int
}
