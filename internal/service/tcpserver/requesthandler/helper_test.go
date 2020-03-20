package requesthandler

import (
	"fmt"
	"testing"

	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"

	"github.com/dickynovanto1103/User-Management-System/internal/model"

	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDataFromCacheNormal(t *testing.T) {
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

func TestGetUserDataFromCacheErrorUsername(t *testing.T) {
	mockCtrl, mockRedis := getRedisErrorGetUsername(t)
	defer mockCtrl.Finish()

	user, err := getUserDataFromCache("user1", mockRedis)
	assert.Equal(t, user, model.User{})
	assert.Equal(t, err, fmt.Errorf("error getting key user1username"))
}

func TestGetUserDataFromCacheErrorNickname(t *testing.T) {
	mockCtrl, mockRedis := getRedisErrorGetNickname(t)
	defer mockCtrl.Finish()

	user, err := getUserDataFromCache("user1", mockRedis)
	assert.Equal(t, user, model.User{})
	assert.Equal(t, err, fmt.Errorf("error getting key user1nickname"))
}

func TestGetUserDataFromCacheErrorProfile(t *testing.T) {
	mockCtrl, mockRedis := getRedisErrorGetProfile(t)
	defer mockCtrl.Finish()

	user, err := getUserDataFromCache("user1", mockRedis)
	assert.Equal(t, user, model.User{})
	assert.Equal(t, err, fmt.Errorf("error getting key user1profile"))
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

func getRedisErrorGetUsername(t *testing.T) (*gomock.Controller, redis.Redis) {
	mockCtrl := gomock.NewController(t)
	mockRedis := redis.NewMockRedis(mockCtrl)

	key := "user1username"
	mockRedis.EXPECT().Get(key).Return("", fmt.Errorf("error getting key %v", key))

	return mockCtrl, mockRedis
}

func getRedisErrorGetNickname(t *testing.T) (*gomock.Controller, redis.Redis) {
	mockCtrl := gomock.NewController(t)
	mockRedis := redis.NewMockRedis(mockCtrl)

	key := "user1username"
	mockRedis.EXPECT().Get(key).Return("", nil)
	key = "user1nickname"
	mockRedis.EXPECT().Get(key).Return("", fmt.Errorf("error getting key %v", key))

	return mockCtrl, mockRedis
}

func getRedisErrorGetProfile(t *testing.T) (*gomock.Controller, redis.Redis) {
	mockCtrl := gomock.NewController(t)
	mockRedis := redis.NewMockRedis(mockCtrl)

	key := "user1username"
	mockRedis.EXPECT().Get(key).Return("", nil)
	key = "user1nickname"
	mockRedis.EXPECT().Get(key).Return("", nil)
	key = "user1profile"
	mockRedis.EXPECT().Get(key).Return("", fmt.Errorf("error getting key %v", key))

	return mockCtrl, mockRedis
}

func TestGetUserDataCacheSuccessful(t *testing.T) {
	mockRedisCtrl, mockRedis := getRedisNormal(t)
	mockDBCtrl, mockDB := getDBMock(t)
	defer mockRedisCtrl.Finish()
	defer mockDBCtrl.Finish()

	username := "user1"
	user, err := getUserData(username, mockRedis, mockDB)
	expectedUser := model.User{
		Username:   username,
		Nickname:   username,
		ProfileURL: "default",
	}
	assert.Equal(t, user, expectedUser)
	assert.Equal(t, err, nil)
}

func getDBMock(t *testing.T) (*gomock.Controller, dbsql.DB) {
	mockCtrl := gomock.NewController(t)
	dbMock := dbsql.NewMockDB(mockCtrl)

	return mockCtrl, dbMock
}

func getDBMockNormal(t *testing.T) (*gomock.Controller, dbsql.DB) {
	mockCtrl := gomock.NewController(t)
	dbMock := dbsql.NewMockDB(mockCtrl)

	username := "user1"
	user := model.User{
		Username:   username,
		Nickname:   username,
		ProfileURL: "default",
	}
	dbMock.EXPECT().GetUser(username).Return(user, nil)

	return mockCtrl, dbMock
}
