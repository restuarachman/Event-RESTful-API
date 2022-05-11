package mysql

import (
	"ticketing/model/domain"

	"gorm.io/gorm"
)

type DBOrderService struct {
	db *gorm.DB
}

func NewDBOrderService(db *gorm.DB) DBOrderService {
	return DBOrderService{
		db: db,
	}
}

func (os DBOrderService) Save(order domain.Order) (domain.Order, error) {
	tx := os.db.Save(&order)
	err := tx.Error
	return order, err
}
func (os DBOrderService) GetAll() ([]domain.Order, error) {
	orders := []domain.Order{}
	tx := os.db.Find(&orders)
	err := tx.Error
	return orders, err
}
func (os DBOrderService) Get(id uint) (domain.Order, error) {
	order := domain.Order{}
	tx := os.db.Find(&order, id)
	err := tx.Error
	return order, err
}
func (os DBOrderService) Delete(order domain.Order) (domain.Order, error) {
	tx := os.db.Delete(&order)
	err := tx.Error
	return order, err
}

func (os DBOrderService) GetAllByUser(user_id uint) ([]domain.Order, error) {
	orders := []domain.Order{}
	tx := os.db.Find(&orders, "user_id = ?", user_id)
	err := tx.Error
	return orders, err
}
