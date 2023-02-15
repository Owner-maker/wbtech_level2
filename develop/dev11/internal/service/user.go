package service

import (
	"dev11/internal/models"
	"dev11/internal/repository/cache"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	hoursDay         = 24
	daysWeek         = 7
	averageDaysMonth = 30
)

func (s *Service) GetEventsForDay(userId string, date time.Time) ([]models.Event, error) {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return nil, err
	}
	events := make([]models.Event, 0)
	for _, v := range user.Events {
		if date.Truncate(hoursDay * time.Hour).Equal(v.Date.Truncate(hoursDay * time.Hour)) {
			events = append(events, v)
		}
	}
	return events, nil
}

func (s *Service) GetEventsForWeek(userId string, startWeekDate time.Time) ([]models.Event, error) {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return nil, err
	}
	events := make([]models.Event, 0)
	var difTime time.Duration

	for _, v := range user.Events {
		difTime = v.Date.Sub(startWeekDate)
		if difTime > 0 && difTime < time.Hour*hoursDay*daysWeek {
			events = append(events, v)
		}
	}
	return events, nil
}

func (s *Service) GetEventsForMonth(userId string, startMonthDate time.Time) ([]models.Event, error) {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return nil, err
	}
	events := make([]models.Event, 0)
	var difTime time.Duration

	for _, v := range user.Events {
		difTime = v.Date.Sub(startMonthDate)
		if difTime > 0 && difTime < time.Hour*hoursDay*averageDaysMonth {
			events = append(events, v)
		}
	}
	return events, nil
}

func (s *Service) CreateEvent(userId string, event models.Event) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	eventId := strconv.Itoa(len(user.Events) + 1)
	event.Id = eventId
	user.Events[eventId] = event
	return nil
}

func (s *Service) UpdateEvent(userId string, event models.Event) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	if _, found := user.Events[event.Id]; !found {
		return cache.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user id = %s", event.Id, userId),
			http.StatusBadRequest)
	}
	user.Events[event.Id] = event
	return nil
}

func (s *Service) DeleteEvent(userId, eventId string) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	if _, found := user.Events[eventId]; !found {
		return cache.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user id = %s", eventId, userId),
			http.StatusBadRequest)
	}
	delete(user.Events, eventId)
	return nil
}
