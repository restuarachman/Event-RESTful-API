package mysql

import (
	"ticketing/model/domain"

	"gorm.io/gorm"
)

type DBUserService struct {
	db *gorm.DB
}

func NewDBUserService(db *gorm.DB) *DBUserService {
	return &DBUserService{
		db: db,
	}
}

func (us *DBUserService) Save(user domain.User) (domain.User, error) {
	tx := us.db.Save(&user)
	err := tx.Error
	return user, err
}

func (us *DBUserService) GetAll() ([]domain.User, error) {
	users := []domain.User{}
	tx := us.db.Find(&users)
	err := tx.Error
	return users, err
}

func (us *DBUserService) Get(id uint) (domain.User, error) {
	user := domain.User{}
	tx := us.db.Find(&user, id)
	err := tx.Error
	return user, err
}

func (us *DBUserService) Delete(user domain.User) (domain.User, error) {
	tx := us.db.Delete(&user)
	err := tx.Error
	return user, err
}

func (us *DBUserService) Login(user domain.User) (domain.User, error) {
	tx := us.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	err := tx.Error
	return user, err
}
