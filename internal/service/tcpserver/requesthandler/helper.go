package requesthandler

import (
	"log"
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"

	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpserver/responsehandler"
)

func getUserData(username string, redis redis.Redis, db dbsql.DB) (model.User, error) {
	data, err := getUserDataFromCache(username, redis)
	if err != nil {
		log.Println("fail to get data from cache", err)
		data, err = db.GetUser(username)
		if err != nil {
			log.Println("fail to get user data from database", err)
			return model.User{}, err
		}
		log.Println("data from database: ", data)
		setUserDataCache(data, redis)
	}
	return data, err
}

func getUserDataFromCache(username string, redis redis.Redis) (model.User, error) {
	var data model.User
	var err error
	data.Username, err = redis.Get(username + model.CodeUsername)
	if err != nil {
		return model.User{}, err
	}
	data.Nickname, err = redis.Get(username + model.CodeNickname)
	if err != nil {
		return model.User{}, err
	}
	data.ProfileURL, err = redis.Get(username + model.CodeProfile)
	if err != nil {
		return model.User{}, err
	}
	return data, nil
}

func setUserDataCache(data model.User, redis redis.Redis) {
	redis.Set(data.Username+model.CodeUsername, data.Username, 5*time.Minute)
	redis.Set(data.Username+model.CodeNickname, data.Nickname, 5*time.Minute)
	redis.Set(data.Username+model.CodeProfile, data.ProfileURL, 5*time.Minute)
}

func sendResponseBack(username string, err error, redis redis.Redis, db dbsql.DB) model.Response {
	if err != nil {
		log.Println("error update statement", err)
		return responsehandler.ResponseForbidden()
	}

	data, err := getUserData(username, redis, db)
	if err != nil {
		log.Println("error getting user data from database: ", err)
		return responsehandler.ResponseForbidden()
	}

	return responsehandler.ResponseUser(data)
}
