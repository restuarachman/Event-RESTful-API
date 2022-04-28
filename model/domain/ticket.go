package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	name        string
	description string
	price       int
}
