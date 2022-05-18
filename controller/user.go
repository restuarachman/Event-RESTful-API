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

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	return NewSuccessResponse(c, response.ToUserResponse(user))
}

func (uc UserController) Update(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	user := domain.User{}

	c.Bind(&user)

	jwtID, _ := mid.ExtractTokenUser(c)

	user, err := uc.us.Update(uint(user_id), user, jwtID)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToUserResponse(user))
}

func (uc UserController) Delete(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))

	jwtID, _ := mid.ExtractTokenUser(c)
	user, err := uc.us.Delete(uint(user_id), jwtID)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
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

// test appleyboy
func (uc UserController) Register(c echo.Context) error {
	var user domain.User
	c.Bind(&user)

	user, err := uc.us.Add(user)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToUserResponse(user))
}
