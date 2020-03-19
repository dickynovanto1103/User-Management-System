package requesthandler

import (
	"testing"

	"github.com/dickynovanto1103/User-Management-System/internal/model"

	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDataFromCache(t *testing.T) {
	mockCtrl, mockRedis := getRedisNormal(t)
	defer mockCtrl.Finish()

	user, err := getUserDataFromCache("user1", mockRedis)
	assert.Equal(t, err, nil)
	expectedUser := model.User{
		Username:   "user1",
		Nickname:   "user1",
		ProfileURL: "default",
	}
	assert.Equal(t, user, expectedUser)

}

func getRedisNormal(t *testing.T) (*gomock.Controller, redis.Redis) {
	mockCtrl := gomock.NewController(t)
	mockRedis := redis.NewMockRedis(mockCtrl)

	key := "user1username"
	mockRedis.EXPECT().Get(key).Return("user1", nil)
	key = "user1nickname"
	mockRedis.EXPECT().Get(key).Return("user1", nil)
	key = "user1profile"
	mockRedis.EXPECT().Get(key).Return("default", nil)
	return mockCtrl, mockRedis
}
