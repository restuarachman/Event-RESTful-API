package response

import "ticketing/model/domain"

type UserResponse struct {
	ID       uint
	Username string
	Name     string
	Email    string
	Phone    string
	Role     string
}
type UserLoginResponse struct {
	ID       uint
	Username string
	Token    string
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}
}

func ToUserLoginResponse(user domain.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}
}

func ToUserListResponse(users []domain.User) []UserResponse {
	var response []UserResponse
	for _, value := range users {
		response = append(response, ToUserResponse(value))
	}
	return response
}
