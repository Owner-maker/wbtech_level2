package models

const initialEventsMapSize = 10

type User struct {
	Id     string           `json:"id"`
	Events map[string]Event `json:"events"`
}

func NewUser(id string) User {
	return User{
		Id:     id,
		Events: make(map[string]Event, initialEventsMapSize),
	}
}
