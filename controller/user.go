package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing/helper/encrypt"
	mid "ticketing/middleware"
	"ticketing/model/domain"
	"ticketing/model/response"
	"ticketing/model/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	us service.UserService
}

func NewUserController(us service.UserService) UserController {
	return UserController{
		us: us,
	}
}

func (uc UserController) GetAll(c echo.Context) error {
	users, err := uc.us.GetAll()
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToUserListResponse(users))
}

func (uc UserController) Get(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	user, err := uc.us.Get(uint(user_id))

	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if user.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("user not found"))
	}
	return NewSuccessResponse(c, response.ToUserResponse(user))
}

func (uc UserController) Update(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	user, err := uc.us.Get(uint(user_id))

	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if user.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("user not found"))
	}
	c.Bind(&user)

	hashPassword, err := encrypt.Hash(user.Password)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	user.Password = hashPassword

	user, err = uc.us.Save(user)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToUserResponse(user))
}

func (uc UserController) Delete(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	user, err := uc.us.Get(uint(user_id))

	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if user.ID == 0 {
		return NewErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("user not found"))
	}
	user, err = uc.us.Delete(user)
	return NewSuccessResponse(c, response.ToUserResponse(user))
}

func (uc UserController) Login(c echo.Context) error {
	var user domain.User

	c.Bind(&user)
	userDB, err := uc.us.GetByUsername(user)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if !encrypt.ValidateHash(user.Password, userDB.Password) {
		return NewErrorResponse(c, http.StatusForbidden, fmt.Errorf("Username or Password invalid"))
	}

	token, err := mid.CreateToken(userDB.ID, userDB.Role)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToUserLoginResponse(userDB, token))
}

func (uc UserController) Register(c echo.Context) error {
	var user domain.User
	c.Bind(&user)

	hashPassword, err := encrypt.Hash(user.Password)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	user.Password = hashPassword

	user, err = uc.us.Save(user)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToUserResponse(user))
}
