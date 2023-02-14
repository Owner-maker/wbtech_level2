package models

import "time"

type Event struct {
	Id          string    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}
