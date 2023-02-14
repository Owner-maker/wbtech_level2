package cache

import (
	"dev11/internal/models"
	"errors"
	"fmt"
	"net/http"
)

type UserCacheRepo struct {
	cch *Cache
}

func NewOrderCache(cch *Cache) *UserCacheRepo {
	return &UserCacheRepo{cch: cch}
}

func (o *UserCacheRepo) PutUser(id string, user models.User) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	o.cch.Data[id] = user
}

func (o *UserCacheRepo) PutUsersEvent(userId string, event models.Event) error {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()

	user, err := o.GetUser(userId)
	if err != nil {
		return err
	}

	user.Events[event.Id] = event
	return nil
}

func (o *UserCacheRepo) GetUser(id string) (*models.User, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()

	if userData, found := o.cch.Data[id]; found {
		return &userData, nil
	}
	return nil, NewErrorHandler(
		errors.New(fmt.Sprintf("failed to find user with id = %s", id)),
		http.StatusBadRequest)
}
