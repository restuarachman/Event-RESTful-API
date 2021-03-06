package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username" form:"username"  gorm:"not null;unique"`
	Password string  `json:"password" form:"password"`
	Name     string  `json:"name" form:"name"`
	Email    string  `json:"email" form:"email"`
	Phone    string  `json:"phone" form:"phone"`
	Role     string  `json:"role" form:"role"`
	Events   []Event `gorm:"foreignKey:UserId"`
	Orders   []Order `gorm:"foreignKey:UserId"`
}
