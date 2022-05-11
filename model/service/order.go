package service

import "ticketing/model/domain"

type OrderService interface {
	Save(order domain.Order) (domain.Order, error)
	GetAll() ([]domain.Order, error)
	Get(id uint) (domain.Order, error)
	Delete(order domain.Order) (domain.Order, error)
	GetAllByUser(user_id uint) ([]domain.Order, error)
}
