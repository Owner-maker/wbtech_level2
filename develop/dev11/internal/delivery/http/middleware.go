package http

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

func (h *Handler) decodeCreateEventBodyJSON(r *http.Request) (*createEventInput, error) {
	input := &createEventInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (h *Handler) decodeUpdateEventBodyJSON(r *http.Request) (*updateEventInput, error) {
	input := &updateEventInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (h *Handler) decodeDeleteEventBodyJSON(r *http.Request) (*deleteEventInput, error) {
	input := &deleteEventInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func getParamsInput(url *url.URL) (string, time.Time, error) {
	userId := url.Query().Get("user_id")
	date := url.Query().Get("date")

	inpDate := InputDate{}
	err := inpDate.UnmarshalJSON([]byte(date))
	if err != nil {
		return "", time.Time{}, err
	}

	return userId, time.Time(inpDate), nil
}
