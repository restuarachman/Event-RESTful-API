package mysql

import (
	"errors"
	"ticketing/helper/encrypt"
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

func (us *DBUserService) Add(user domain.User) (domain.User, error) {
	pw, _ := encrypt.Hash(user.Password)
	user.Password = pw
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
	if user.ID == 0 {
		return user, errors.New("User not found")
	}
	return user, err
}

func (us *DBUserService) Update(id uint, user domain.User, jwtID uint) (domain.User, error) {
	userDB, err := us.Get(id)
	if err != nil {
		return user, err
	}
	if userDB.ID != jwtID {
		return domain.User{}, errors.New("forbidden")
	}

	if user.Name != "" {
		userDB.Name = user.Name
	}
	if user.Email != "" {
		userDB.Email = user.Email
	}
	if user.Password != "" {
		pw, _ := encrypt.Hash(user.Password)
		userDB.Password = pw
	}
	if user.Phone != "" {
		userDB.Phone = user.Phone
	}

	tx := us.db.Save(&userDB)
	err = tx.Error
	return userDB, err
}

func (us *DBUserService) Delete(id uint, jwtID uint) (domain.User, error) {
	user, err := us.Get(id)
	if err != nil {
		return user, err
	}
	if id != jwtID {
		return domain.User{}, errors.New("forbidden")
	}
	tx := us.db.Delete(&user)
	err = tx.Error
	return user, err
}

func (us *DBUserService) GetByUsername(user domain.User) (domain.User, error) {
	tx := us.db.Find(&user, "username = ?", user.Username)
	err := tx.Error
	return user, err
}
