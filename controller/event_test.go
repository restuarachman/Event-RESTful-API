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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GetEventsResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data []domain.Event
}

type GetEventResponse struct {
	Meta struct {
		Status  int
		Message string
	}
	Data domain.Event
}

func TestEventControllerCreate(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Add(domain.User{Username: "restu", Password: "rahasia", Name: "Restu Aditya R"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned event failed")
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

	// create
	newEventJson, _ := json.Marshal(map[string]interface{}{
		"name":        "Ciater",
		"description": "Pemandian air panas",
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newEventJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	es := mock.NewMockEventService()
	ec := NewEventController(es)

	if err := middleware.JWT([]byte(constants.SECRET_JWT))(ec.Create)(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	events, _ := es.GetAll()
	if len(events) != 1 {
		t.Errorf("Expecting len(users) to be 1, get %d", len(events))
	}

	if events[0].Name != "Ciater" {
		t.Errorf("Expectiong name to be Ciater, get %s", events[0].Name)
	}
}

func TestEventControllerGetAll(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	es := mock.NewMockEventService()
	es.Add(domain.Event{Name: "Ciater"}, 1)
	es.Add(domain.Event{Name: "Dunia Fantasi"}, 2)
	es.Add(domain.Event{Name: "Rumah Hobit"}, 3)

	ec := NewEventController(es)

	if err := ec.GetAll(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetEventsResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned event failed")
	}
	events := response.Data

	if len(events) != 3 {
		t.Errorf("expecting len(users) is 3, get %d", len(events))
	}
	if events[0].Name != "Ciater" {
		t.Errorf("expection events[0].Name is Ciater, get %s", events[0].Name)
	}
	if events[1].Name != "Dunia Fantasi" {
		t.Errorf("expection events[1].Name is Dunia Fantasi, get %s", events[1].Name)
	}
	if events[2].Name != "Rumah Hobit" {
		t.Errorf("expection events[2].Name is Rumah Hobit, get %s", events[2].Name)
	}
}

func TestEventControllerGet(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:event_id")
	c.SetParamNames("event_id")
	c.SetParamValues("1")

	es := mock.NewMockEventService()
	es.Add(domain.Event{Name: "Ciater"}, 1)
	es.Add(domain.Event{Name: "Dunia Fantasi"}, 2)
	es.Add(domain.Event{Name: "Rumah Hobit"}, 3)

	ec := NewEventController(es)

	if err := ec.Get(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetEventResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned event failed")
	}
	event := response.Data

	if event.Name != "Ciater" {
		t.Errorf("expection event.Name is Rumah Hobit, get %s", event.Name)
	}
	if event.UserId != 1 {
		t.Errorf("expection event.UserId is 1, get %d", event.UserId)
	}
}

func TestEventControllerUpdate(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Add(domain.User{Username: "restu", Password: "rahasia", Name: "Restu Aditya R"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned event failed")
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

	// update
	newEventJson, _ := json.Marshal(map[string]interface{}{
		"name":        "Ciater",
		"description": "Pemandian air panas",
	})

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(newEventJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:event_id")
	c.SetParamNames("event_id")
	c.SetParamValues("1")

	es := mock.NewMockEventService()
	es.Add(domain.Event{Name: "Ciater"}, 1)
	es.Add(domain.Event{Name: "Dunia Fantasi"}, 2)
	es.Add(domain.Event{Name: "Rumah Hobit"}, 3)

	ec := NewEventController(es)

	if err := middleware.JWT([]byte(constants.SECRET_JWT))(ec.Update)(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetEventResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned event failed")
	}
	event := response.Data

	if event.Name != "Ciater" {
		t.Errorf("expection event.Name is Rumah Hobit, get %s", event.Name)
	}
	if event.UserId != 1 {
		t.Errorf("expection event.UserId is 1, get %d", event.UserId)
	}
}

func TestEventControllerDelete(t *testing.T) {
	e := echo.New()

	us := mock.NewMockUserService()
	us.Add(domain.User{Username: "restu", Password: "rahasia", Name: "Restu Aditya R"})

	// Login First
	loginUser := domain.User{Username: "restu", Password: "rahasia"}

	loginInfo, err := json.Marshal(loginUser)
	if err != nil {
		t.Error("Marhalling returned event failed")
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

	// update
	newEventJson, _ := json.Marshal(map[string]interface{}{
		"name":        "Ciater",
		"description": "Pemandian air panas",
	})

	req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(newEventJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", userLogin.Data.Token))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:event_id")
	c.SetParamNames("event_id")
	c.SetParamValues("1")

	es := mock.NewMockEventService()
	es.Add(domain.Event{Name: "Ciater"}, 1)
	es.Add(domain.Event{Name: "Dunia Fantasi"}, 2)
	es.Add(domain.Event{Name: "Rumah Hobit"}, 3)

	ec := NewEventController(es)

	if err := middleware.JWT([]byte(constants.SECRET_JWT))(ec.Delete)(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetEventResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned event failed")
	}

	if response.Meta.Status != 200 {
		t.Errorf("expection code is 200, get %d", response.Meta.Status)
	}
	if response.Data.ID != 1 {
		t.Errorf("expection response.Data.ID is 1, get %d", response.Data.ID)
	}
}

func TestEventControllerGetAllEventByUserId(t *testing.T) {
	e := echo.New()

	es := mock.NewMockEventService()
	ec := NewEventController(es)

	es.Add(domain.Event{Name: "Ciater"}, 1)
	es.Add(domain.Event{Name: "Dunia Fantasi"}, 1)
	es.Add(domain.Event{Name: "Rumah Hobit"}, 2)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// case 1 success
	c.SetPath("/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("1")

	if err := ec.GetAllEventByUserId(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
	}

	response := GetEventsResponse{}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Error("unmarhalling returned event failed")
	}
	events := response.Data

	if len(events) != 2 {
		t.Errorf("expecting len(events) is 3, get %d", len(events))
	}
	if events[0].Name != "Ciater" {
		t.Errorf("expection events[0].Name is Ciater, get %s", events[0].Name)
	}
	if events[1].Name != "Dunia Fantasi" {
		t.Errorf("expection events[1].Name is Dunia Fantasi, get %s", events[1].Name)
	}
}
