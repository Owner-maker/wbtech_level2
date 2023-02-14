package service

import (
	"dev11/internal/models"
	"dev11/internal/repository/cache"
	"time"
)

type User interface {
	GetEventsForDay(id string, date time.Time) ([]models.Event, error)
	GetEventsForWeek(id string, startWeekDate time.Time) ([]models.Event, error)
	GetEventsForMonth(id string, startMonthDate time.Time) ([]models.Event, error)
	CreateEvent(userId string, event models.Event) error
	UpdateEvent(userId string, event models.Event) error
	DeleteEvent(userId, eventId string) error
}

type Service struct {
	cache cache.UserCacheRepo
}

func NewService() *Service {
	return &Service{cache: *cache.NewUserCache(cache.NewCache())}
}
