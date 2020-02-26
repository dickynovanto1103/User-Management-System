package requesthandler

import (
	"log"

	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	responsehandler "github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/responseHandler"
)

type UpdateProfileHandler struct{}

func (handler *UpdateProfileHandler) HandleRequest(mapper map[string]string, redis redis.Redis) model.Response {
	profile := mapper[model.CodeProfile]
	userIDFromCookie := mapper[model.CodeCookie]
	username, err := redis.Get(userIDFromCookie)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		return responsehandler.ResponseForbidden()
	}
	err = UpdateProfile(profile, username)
	if err != nil {
		log.Println(dbsql.ErrorUpdateProfile + " " + err.Error())
		return responsehandler.ResponseError(err)
	}
	user, err := GetUser(username)
	setUserDataCache(user, redis)

	return sendResponseBack(username, err, redis)
}
