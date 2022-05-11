package service

import "ticketing/model/domain"

type OrderDetailsService interface {
	Save(order []domain.OrderDetail) ([]domain.OrderDetail, error)
	GetAll() ([]domain.OrderDetail, error)
	GetByOrderId(order_id uint) ([]domain.OrderDetail, error)
}
