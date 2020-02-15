package main

import (
	"github.com/dickynovanto1103/User-Management-System/internal/request"
	"github.com/dickynovanto1103/User-Management-System/tcpserver/requestHandler"
)

var mapperReqIdToCommand = map[int]requestHandler.RequestHandler {
	request.RequestLogin: &requestHandler.AuthenticationHandler{},
	request.RequestUserInfo: &requestHandler.InfoHandler{},
	request.RequestUpdateNickname: &requestHandler.UpdateNicknameHandler{},
	request.RequestUpdateProfile: &requestHandler.UpdateProfileHandler{},
}
