package cache

import (
	"dev11/internal/models"
	"fmt"
	"net/http"
)

type UserCacheRepo struct {
	cch *Cache
}

func NewUserCache(cch *Cache) *UserCacheRepo {
	c := UserCacheRepo{cch: cch}
	c.addTestUser()
	return &c
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
		fmt.Errorf("failed to find user with id = %s", id),
		http.StatusBadRequest)
}
