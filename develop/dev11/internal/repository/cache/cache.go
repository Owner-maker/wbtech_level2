package cache

import (
	"dev11/internal/models"
	"sync"
)

const initialMapSize = 10

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string]models.User
}

func NewCache() *Cache {
	var cache Cache
	cache.Data = make(map[string]models.User, initialMapSize)
	return &cache
}
