package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.DB
	name        string
	description string
	date_start  time.Time
	date_end    time.Time
	time        string
	location    string
}
