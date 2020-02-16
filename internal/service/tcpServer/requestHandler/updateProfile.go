package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer/responseHandler"
	"log"
)

type UpdateProfileHandler struct{}

func (handler *UpdateProfileHandler) HandleRequest(mapper map[string]string) model.Response {
	profile := mapper[model.CodeProfile]
	userIDFromCookie := mapper[model.CodeCookie]
	username, err := redisutil.Get(userIDFromCookie)

	if err != nil {
		log.Println("error when getting user from cookie ", err)
		return responseHandler.ResponseForbidden()
	}
	err = dbutil.UpdateProfile(profile, username)
	if err != nil {
		log.Println(dbutil.ErrorUpdateProfile + " " + err.Error())
		return responseHandler.ResponseError(err)
	}
	user, err := dbutil.GetUser(username)
	setUserDataCache(user)

	return sendResponseBack(username, err)
}
