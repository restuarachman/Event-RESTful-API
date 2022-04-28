package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	user_id       int
	checkout_date string
	status        string
	total_price   int
}
