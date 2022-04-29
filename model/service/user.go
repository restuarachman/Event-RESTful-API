package service

import "ticketing/model/domain"

type UserService interface {
	Save(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	Get(id uint) (domain.User, error)
	Delete(user domain.User) (domain.User, error)
	GetByUsername(user domain.User) (domain.User, error)
}
