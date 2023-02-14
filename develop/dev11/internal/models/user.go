package models

const initialEventsMapSize = 10

type User struct {
	Id     string           `json:"id" validate:"required"`
	Events map[string]Event `json:"events"`
}

func newUser(id string) User {
	return User{
		Id:     id,
		Events: make(map[string]Event, initialEventsMapSize),
	}
}

func newUserWithEvents(id string, events map[string]Event) User {
	return User{
		Id:     id,
		Events: events,
	}
}
