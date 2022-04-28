package mock

import (
	"ticketing/model/domain"
)

type MockUserService struct {
	data []domain.User
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		data: []domain.User{},
	}
}

// Save(user model.User) (model.User, error)
// GetAll() ([]model.User, error)
// Get(id uint) (model.User, error)
// Delete(user model.User) (model.User, error)
// Login(user model.User) (model.User, error)
