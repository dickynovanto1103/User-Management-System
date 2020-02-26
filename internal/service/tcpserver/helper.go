package tcpserver

import (
	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpserver/requesthandler"
)

var mapperReqIdToCommand = map[int]requesthandler.RequestHandler{
	model.RequestLogin:          &requesthandler.AuthenticationHandler{},
	model.RequestUserInfo:       &requesthandler.InfoHandler{},
	model.RequestUpdateNickname: &requesthandler.UpdateNicknameHandler{},
	model.RequestUpdateProfile:  &requesthandler.UpdateProfileHandler{},
}

func HandleRequest(requestID int32, mapper map[string]string) model.Response {
	command := mapperReqIdToCommand[int(requestID)]
	return command.HandleRequest(mapper, container.RedisImpl)
}
