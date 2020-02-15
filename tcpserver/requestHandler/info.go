package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
	"github.com/dickynovanto1103/User-Management-System/tcpserver/responseHandler"
	"log"
)

type InfoHandler struct{}

func (handler *InfoHandler) HandleRequest(mapper map[string]string) response.Response {
	userIDFromCookie := mapper[user.CodeCookie]
	username, err := redisutil.Get(userIDFromCookie)
	if err != nil {
		log.Println("error retrieving user data from cookie by redis: ", err)
		return responseHandler.ResponseForbidden()
	}
	data, err := getUserData(username)
	if err != nil {
		log.Println("fail to get user data", err)
		return responseHandler.ResponseForbidden()
	} else {
		return responseHandler.ResponseUser(data)
	}
}
