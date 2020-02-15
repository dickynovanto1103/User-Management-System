package responseHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
)

func ResponseForbidden() response.Response{
	var mapper = make(map[string]string)
	mapper[response.ResponseCode] = response.ResponseKeyForbidden
	return response.Response{ResponseID: response.ResponseForbidden, Data: mapper}
}

func ResponseUser(data user.User) response.Response{
	var mapper = make(map[string]string)
	mapper = getMapFromUser(data)
	return response.Response{ResponseID: response.ResponseOK, Data: mapper}
}

func ResponseError(err error) response.Response {
	var mapper = make(map[string]string)
	mapper[response.ResponseCode] = response.ResponseKeyError
	mapper[response.ResponseKeyError] = err.Error()
	return response.Response{ResponseID: response.ResponseError, Data: mapper}
}


func getMapFromUser(data user.User) map[string]string {
	var mapper = make(map[string]string)
	mapper[user.CodeUsername] = data.Username
	mapper[user.CodeNickname] = data.Nickname
	mapper[user.CodeProfile] = data.ProfileURL
	return mapper
}