package main

import (
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
)

func responseForbidden() response.Response{
	var mapper = make(map[string]string)
	mapper[response.ResponseCode] = response.ResponseKeyForbidden
	return response.Response{ResponseID: response.ResponseForbidden, Data: mapper}
}

func responseUser(data user.User) response.Response{
	var mapper = make(map[string]string)
	mapper = getMapFromUser(data)
	return response.Response{ResponseID: response.ResponseOK, Data: mapper}
}

func responseError(err error) response.Response {
	var mapper = make(map[string]string)
	mapper[response.ResponseCode] = response.ResponseKeyError
	mapper[response.ResponseKeyError] = err.Error()
	return response.Response{ResponseID: response.ResponseError, Data: mapper}
}