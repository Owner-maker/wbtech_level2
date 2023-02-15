package http

import (
	"dev11/internal/models"
	"encoding/json"
	"strings"
	"time"
)

type InputDate time.Time

type createEventInput struct {
	UserId      string    `json:"userId" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Date        InputDate `json:"date" validate:"required"`
}

type updateEventInput struct {
	UserId      string    `json:"userId" validate:"required"`
	EventId     string    `json:"eventId" validate:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        InputDate `json:"date"`
}

type deleteEventInput struct {
	UserId  string `json:"userId" validate:"required"`
	EventId string `json:"eventId" validate:"required"`
}

type successEventOutput struct {
	Result string `json:"result"`
}

type eventsOutput struct {
	Result []models.Event `json:"result"`
}

type errorOutput struct {
	Error string `json:"error"`
}

func newSuccessEventOutput(result string) successEventOutput {
	return successEventOutput{Result: result}
}

func newEventsOutput(result []models.Event) eventsOutput {
	return eventsOutput{Result: result}
}

func newErrorOutput(message string) errorOutput {
	return errorOutput{Error: message}
}

func (i *InputDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*i = InputDate(t)
	return nil
}

func (i InputDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(i))
}
