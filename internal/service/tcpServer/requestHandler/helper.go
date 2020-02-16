package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/responseHandler"
	"log"
	"time"
)

func getUserData(username string) (model.User, error) {
	data, err := getUserDataFromCache(username)
	if err != nil {
		log.Println("fail to get data from cache", err)
		data, err = dbutil.GetUser(username)
		if err != nil {
			log.Println("fail to get user data from database", err)
			return model.User{}, err
		}
		log.Println("data from database: ", data)
		setUserDataCache(data)
	}
	return data, err
}

func getUserDataFromCache(username string) (model.User, error) {
	var data model.User
	var err error
	data.Username, err = redisutil.Get(username + model.CodeUsername)
	if err != nil {
		return model.User{}, err
	}
	data.Nickname, err = redisutil.Get(username + model.CodeNickname)
	if err != nil {
		return model.User{}, err
	}
	data.ProfileURL, err = redisutil.Get(username + model.CodeProfile)
	if err != nil {
		return model.User{}, err
	}
	return data, nil
}

func setUserDataCache(data model.User) {
	redisutil.Set(data.Username+model.CodeUsername, data.Username, 5*time.Minute)
	redisutil.Set(data.Username+model.CodeNickname, data.Nickname, 5*time.Minute)
	redisutil.Set(data.Username+model.CodeProfile, data.ProfileURL, 5*time.Minute)
}

func sendResponseBack(username string, err error) model.Response {
	if err != nil {
		log.Println("error update statement", err)
		return responseHandler.ResponseForbidden()
	} else {
		data, err := getUserData(username)
		if err != nil {
			log.Println("error getting user data from database: ", err)
			return responseHandler.ResponseForbidden()
		}
		return responseHandler.ResponseUser(data)
	}
}