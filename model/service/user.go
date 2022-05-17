package service

import "ticketing/model/domain"

type UserService interface {
	Add(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	Get(id uint) (domain.User, error)
	Update(id uint, user domain.User, jwtID uint) (domain.User, error)
	Delete(id uint, jwtID uint) (domain.User, error)
	GetByUsername(user domain.User) (domain.User, error)
}
