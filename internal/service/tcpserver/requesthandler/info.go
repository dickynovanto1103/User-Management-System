package requesthandler

import (
	"log"

	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	responsehandler "github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/responseHandler"
)

type InfoHandler struct{}

func (handler *InfoHandler) HandleRequest(mapper map[string]string, redis redis.Redis) model.Response {
	userIDFromCookie := mapper[model.CodeCookie]
	username, err := redis.Get(userIDFromCookie)
	if err != nil {
		log.Println("error retrieving user data from cookie by redis: ", err)
		return responsehandler.ResponseForbidden()
	}
	data, err := getUserData(username, redis)
	if err != nil {
		log.Println("fail to get user data", err)
		return responsehandler.ResponseForbidden()
	}

	return responsehandler.ResponseUser(data)
}
