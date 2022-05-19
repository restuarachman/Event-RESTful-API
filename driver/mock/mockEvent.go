package mock

import (
	"errors"
	"ticketing/model/domain"
)

type MockEventService struct {
	data  []domain.Event
	numid int
}

func NewMockEventService() *MockEventService {
	return &MockEventService{
		data: []domain.Event{},
	}
}

func (ec *MockEventService) Add(event domain.Event, user_id uint) (domain.Event, error) {
	ec.numid++
	event.ID = uint(ec.numid)
	event.UserId = user_id

	ec.data = append(ec.data, event)
	return event, nil
}

func (ec *MockEventService) GetAll() ([]domain.Event, error) {
	return ec.data, nil
}

func (ec *MockEventService) Get(id uint) (domain.Event, error) {
	for _, val := range ec.data {
		if val.ID == id {
			return val, nil
		}
	}
	return domain.Event{}, errors.New("Event not found")
}

func (ec *MockEventService) Update(event_id uint, event domain.Event, jwtID uint) (domain.Event, error) {
	event, err := ec.Get(event_id)
	if err != nil {
		return domain.Event{}, err
	}
	if event.ID != jwtID {
		return domain.Event{}, errors.New("Event not found")
	}
	for _, val := range ec.data {
		if val.ID == event_id {
			val = event
			val.ID = event_id
			return val, nil
		}
	}
	return domain.Event{}, errors.New("Event not found")
}

func (ec *MockEventService) Delete(event_id uint, jwtID uint) (domain.Event, error) {
	event, err := ec.Get(event_id)
	if err != nil {
		return domain.Event{}, err
	}
	if event.ID != jwtID {
		return domain.Event{}, errors.New("Event not found")
	}
	for i, val := range ec.data {
		if val.ID == event_id {
			j := i
			for j < len(ec.data)-1 {
				ec.data[i] = ec.data[j+1]
				j++
			}
			ec.data = ec.data[:j-1]
			return event, nil
		}
	}
	return domain.Event{}, nil
}

func (ec *MockEventService) GetByUserId(user_id uint) ([]domain.Event, error) {
	events := []domain.Event{}
	for _, val := range ec.data {
		if val.UserId == user_id {
			events = append(events, val)
		}
	}
	return events, nil
}
