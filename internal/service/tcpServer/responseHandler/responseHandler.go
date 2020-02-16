package responseHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/model"
)

func ResponseForbidden() model.Response {
	var mapper = make(map[string]string)
	mapper[model.ResponseCode] = model.ResponseKeyForbidden
	return model.Response{ResponseID: model.ResponseForbidden, Data: mapper}
}

func ResponseUser(data model.User) model.Response {
	var mapper = make(map[string]string)
	mapper = getMapFromUser(data)
	return model.Response{ResponseID: model.ResponseOK, Data: mapper}
}

func ResponseError(err error) model.Response {
	var mapper = make(map[string]string)
	mapper[model.ResponseCode] = model.ResponseKeyError
	mapper[model.ResponseKeyError] = err.Error()
	return model.Response{ResponseID: model.ResponseError, Data: mapper}
}

func getMapFromUser(data model.User) map[string]string {
	var mapper = make(map[string]string)
	mapper[model.CodeUsername] = data.Username
	mapper[model.CodeNickname] = data.Nickname
	mapper[model.CodeProfile] = data.ProfileURL
	return mapper
}