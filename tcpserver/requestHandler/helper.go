package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
	"github.com/dickynovanto1103/User-Management-System/tcpserver/responseHandler"
	"log"
	"time"
)

func getUserData(username string) (user.User, error) {
	data, err := getUserDataFromCache(username)
	if err != nil {
		log.Println("fail to get data from cache", err)
		data, err = dbutil.GetUser(username)
		if err != nil {
			log.Println("fail to get user data from database", err)
			return user.User{}, err
		}
		log.Println("data from database: ", data)
		setUserDataCache(data)
	}
	return data, err
}

func getUserDataFromCache(username string) (user.User, error) {
	var data user.User
	var err error
	data.Username, err = redisutil.Get(username + user.CodeUsername)
	if err != nil {
		return user.User{}, err
	}
	data.Nickname, err = redisutil.Get(username + user.CodeNickname)
	if err != nil {
		return user.User{}, err
	}
	data.ProfileURL, err = redisutil.Get(username + user.CodeProfile)
	if err != nil {
		return user.User{}, err
	}
	return data, nil
}

func setUserDataCache(data user.User) {
	redisutil.Set(data.Username+user.CodeUsername, data.Username, 5*time.Minute)
	redisutil.Set(data.Username+user.CodeNickname, data.Nickname, 5*time.Minute)
	redisutil.Set(data.Username+user.CodeProfile, data.ProfileURL, 5*time.Minute)
}

func sendResponseBack(username string, err error) response.Response {
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