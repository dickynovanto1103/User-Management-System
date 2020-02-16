package tcpServer

import (
	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/requestHandler"
)

var mapperReqIdToCommand = map[int]requestHandler.RequestHandler{
	model.RequestLogin:          &requestHandler.AuthenticationHandler{},
	model.RequestUserInfo:       &requestHandler.InfoHandler{},
	model.RequestUpdateNickname: &requestHandler.UpdateNicknameHandler{},
	model.RequestUpdateProfile:  &requestHandler.UpdateProfileHandler{},
}

func HandleRequest(requestID int32, mapper map[string]string) model.Response {
	command := mapperReqIdToCommand[int(requestID)]
	return command.HandleRequest(mapper, container.RedisImpl)
}
