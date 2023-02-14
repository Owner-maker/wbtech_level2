package service

import (
	"dev11/internal/models"
	"dev11/internal/repository/cache"
	"time"
)

type User interface {
	GetEventsForDay(id string, date time.Time) ([]models.Event, error)
	GetEventsForWeek(id string, date time.Time) ([]models.Event, error)
	GetEventsForMonth(id string, date time.Time) ([]models.Event, error)
	CreateEvent(userId string, event models.Event) error
	UpdateEvent(userId, eventId string) error
	DeleteEvent(userId, eventId string) error
}

type Service struct {
	cache.UserCacheRepo
}
