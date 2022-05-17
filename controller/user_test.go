package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketing/driver/mock"
	"ticketing/middleware/constants"
	"ticketing/model/domain"
	"ticketing/model/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GetUsersResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data []domain.User
}

type GetUserResponse struct {
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
	us.Add(domain.User{Username: "dono", Password: "rahasia"})

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

	users, _ := us.GetAll()
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
	us.Add(domain.User{Email: "restu@gmail.com"})
	us.Add(domain.User{Email: "aditya@gmail.com"})
	us.Add(domain.User{Email: "rachman@gmail.com"})

	uc := NewUserController(us)

	if err := uc.GetAll(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetUsersResponse{}

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
	us.Add(domain.User{Username: "restu", Password: "rahasia"})
	us.Add(domain.User{Username: "aditya", Password: "rahasia"})
	us.Add(domain.User{Username: "rachman", Password: "rahasia"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned user failed")
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	logincontext := e.NewContext(loginReq, loginRec)
	logincontext.SetPath("/login")

	uc := NewUserController(us)
	if err := uc.Login(logincontext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var userLogin LoginResponse
	if err := json.Unmarshal(loginRec.Body.Bytes(), &userLogin); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}

	if userLogin.Meta.Status != 200 {
		t.Errorf("Login fail")
	}
	// Get User
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	// uc := NewUserController(us)
	if err := middleware.JWT([]byte(constants.SECRET_JWT))(uc.Get)(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetUserResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("unmarshalling returned person failed: %s", err)
	}

	user := response.Data
	if user.Username != "restu" {
		t.Errorf("expection username is restu, get %s", user.Username)
	}
}

func TestUserControllerUpdate(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Add(domain.User{Username: "restu", Password: "rahasia", Name: "Restu Aditya R"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned user failed")
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	logincontext := e.NewContext(loginReq, loginRec)
	logincontext.SetPath("/login")

	uc := NewUserController(us)
	if err := uc.Login(logincontext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var userLogin LoginResponse
	if err := json.Unmarshal(loginRec.Body.Bytes(), &userLogin); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}

	if userLogin.Meta.Status != 200 {
		t.Errorf("Login fail")
	}

	// Update
	newUserJson, _ := json.Marshal(map[string]interface{}{
		"name": "Restu Aditya Rachman",
	})

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUserJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	if err := middleware.JWT([]byte(constants.SECRET_JWT))(uc.Update)(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetUserResponse{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("unmarshalling returned person failed: %s", err)
	}

	user := response.Data
	if user.Name != "Restu Aditya Rachman" {
		t.Errorf("expection name is Restu Aditya Rachman, get %s", user.Name)
	}
}

func TestUserControllerDelete(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Add(domain.User{Username: "restu", Password: "rahasia", Name: "Restu Aditya R"})
	us.Add(domain.User{Username: "a", Password: "rahasia", Name: "a"})
	us.Add(domain.User{Username: "b", Password: "rahasia", Name: "b"})
	us.Add(domain.User{Username: "c", Password: "rahasia", Name: "c"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned user failed")
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	logincontext := e.NewContext(loginReq, loginRec)
	logincontext.SetPath("/login")

	uc := NewUserController(us)
	if err := uc.Login(logincontext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var userLogin LoginResponse
	if err := json.Unmarshal(loginRec.Body.Bytes(), &userLogin); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}

	if userLogin.Meta.Status != 200 {
		t.Errorf("Login fail")
	}

	// Update
	newUserJson, _ := json.Marshal(map[string]interface{}{
		"name": "Restu Aditya Rachman",
	})

	// Delete
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newUserJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	if err := middleware.JWT([]byte(constants.SECRET_JWT))(uc.Delete)(c); err != nil {
		t.Errorf("unmarshalling returned person failed: %s", err)
	}

	response := GetUserResponse{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("unmarshalling returned person failed: %s", err)
	}

	if response.Meta.Status != 200 {
		t.Errorf("expection code is 200, get %d", response.Meta.Status)
	}
}
