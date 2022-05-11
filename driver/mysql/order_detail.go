package mysql

import (
	"ticketing/model/domain"

	"gorm.io/gorm"
)

type DBOrderDetailService struct {
	db *gorm.DB
}

func NewDBOrderDetailService(db *gorm.DB) DBOrderDetailService {
	return DBOrderDetailService{
		db: db,
	}
}

func (ods DBOrderDetailService) Save(orderDetail []domain.OrderDetail) ([]domain.OrderDetail, error) {
	tx := ods.db.Save(&orderDetail)
	err := tx.Error
	return orderDetail, err
}
func (ods DBOrderDetailService) GetAll() ([]domain.OrderDetail, error) {
	orderDetails := []domain.OrderDetail{}
	tx := ods.db.Find(&orderDetails)
	err := tx.Error
	return orderDetails, err
}
func (ods DBOrderDetailService) GetByOrderId(order_id uint) ([]domain.OrderDetail, error) {
	orderDetails := []domain.OrderDetail{}
	tx := ods.db.Find(&orderDetails, order_id)
	err := tx.Error
	return orderDetails, err
}
