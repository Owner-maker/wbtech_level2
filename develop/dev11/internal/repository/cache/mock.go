package cache

import "dev11/internal/models"

func (o *UserCacheRepo) addTestUser() {
	testUser := models.NewUser("1")

	o.PutUser(testUser.Id, testUser)
}
