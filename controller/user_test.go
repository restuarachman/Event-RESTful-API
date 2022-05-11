package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketing/driver/mock"
	"ticketing/helper/encrypt"
	"ticketing/model/domain"
	"ticketing/model/response"

	"github.com/labstack/echo/v4"
)

type GetAllResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data []domain.User
}

type GetResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data domain.User
}

type LoginResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data response.UserLoginResponse
}

func TestUserControllerLogin(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	pw, _ := encrypt.Hash("rahasia")
	us.Save(domain.User{Username: "dono", Password: pw})

	user := domain.User{Username: "dono", Password: "rahasia"}

	loginInfo, err := json.Marshal(user)
	if err != nil {
		t.Error("Marhalling returned user failed")
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	uc := NewUserController(us)
	if err := uc.Login(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var userLogin LoginResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &userLogin); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}

	if userLogin.Data.Token == "" {
		t.Errorf("token expected")
	}
}

func TestUserControllerRegister(t *testing.T) {
	e := echo.New()

	newUserJson, _ := json.Marshal(map[string]interface{}{
		"username": "restu",
		"name":     "Restu Aditya Rachman",
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUserJson))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	us := mock.NewMockUserService()
	uc := NewUserController(us)

	if err := uc.Register(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	users, err := us.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("Expecting len(users) to be 1, get %d", len(users))
	}
	if users[0].Username != "restu" {
		t.Errorf("Expectiong username to be restu, get %s", users[0].Username)
	}
}

func TestUserControllerGetAll(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	us := mock.NewMockUserService()
	us.Save(domain.User{Email: "restu@gmail.com"})
	us.Save(domain.User{Email: "aditya@gmail.com"})
	us.Save(domain.User{Email: "rachman@gmail.com"})

	uc := NewUserController(us)
	uc.GetAll(c)

	response := GetAllResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned user failed")
	}
	users := response.Data

	if len(users) != 3 {
		t.Errorf("expecting len(users) is 3, get %d", len(users))
	}
	if users[0].Email != "restu@gmail.com" {
		t.Errorf("expection users[0].Email is restu@gmail.com, get %s", users[0].Email)
	}
	if users[1].Email != "aditya@gmail.com" {
		t.Errorf("expection users[0].Email is aditya@gmail.com, get %s", users[1].Email)
	}
	if users[2].Email != "rachman@gmail.com" {
		t.Errorf("expection users[0].Email is rachman@gmail.com, get %s", users[2].Email)
	}
}

func TestUserControllerGet(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Save(domain.User{Username: "restu"})
	us.Save(domain.User{Username: "aditya"})
	us.Save(domain.User{Username: "rachman"})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	uc := NewUserController(us)
	uc.Get(c)

	response := GetAllResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned user failed")
	}
	users := response.Data

	if len(users) != 3 {
		t.Errorf("expecting len(users) is 3, get %d", len(users))
	}
	if users[0].Email != "restu@gmail.com" {
		t.Errorf("expection users[0].Email is restu@gmail.com, get %s", users[0].Email)
	}
	if users[1].Email != "aditya@gmail.com" {
		t.Errorf("expection users[0].Email is aditya@gmail.com, get %s", users[1].Email)
	}
	if users[2].Email != "rachman@gmail.com" {
		t.Errorf("expection users[0].Email is rachman@gmail.com, get %s", users[2].Email)
	}
}
