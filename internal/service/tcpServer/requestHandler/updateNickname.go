package requestHandler

import (
	"log"

	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/responseHandler"
)

type UpdateNicknameHandler struct{}

func (handler *UpdateNicknameHandler) HandleRequest(mapper map[string]string, redis redis.Redis) model.Response {
	nickname := mapper[model.CodeNickname]
	cookieValue := mapper[model.CodeCookie]
	username, err := redis.Get(cookieValue)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		return responseHandler.ResponseForbidden()
	}

	err = UpdateNickname(nickname, username)
	if err != nil {
		log.Println(dbsql.ErrorUpdateProfile + " " + err.Error())
		return responseHandler.ResponseError(err)
	}
	user, err := GetUser(username)
	setUserDataCache(user, redis)

	return sendResponseBack(username, err, redis)
}
