package mock

import (
	"ticketing/helper/encrypt"
	"ticketing/model/domain"
)

type MockUserService struct {
	data  []domain.User
	numid int
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		data: []domain.User{},
	}
}

func (uc *MockUserService) Save(user domain.User) (domain.User, error) {
	uc.numid++
	user.ID = uint(uc.numid)
	pw, _ := encrypt.Hash(user.Password)
	user.Password = pw
	uc.data = append(uc.data, user)
	return user, nil
}
func (uc *MockUserService) GetAll() ([]domain.User, error) {
	return uc.data, nil
}
func (uc *MockUserService) Get(id uint) (domain.User, error) {
	for _, val := range uc.data {
		if val.ID == id {
			return val, nil
		}
	}
	return domain.User{}, nil
}
func (uc *MockUserService) Delete(user domain.User) (domain.User, error) {

	for i, val := range uc.data {
		if val.ID == user.ID {
			j := i
			for j < len(uc.data)-1 {
				uc.data[i] = uc.data[j+1]
				j++
			}
			uc.data = uc.data[:j-1]
			return user, nil
		}
	}
	return domain.User{}, nil
}
func (uc *MockUserService) GetByUsername(user domain.User) (domain.User, error) {
	for _, val := range uc.data {
		if val.Username == user.Username {
			return val, nil
		}
	}
	return domain.User{}, nil
}
