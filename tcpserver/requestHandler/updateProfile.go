package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
	"github.com/dickynovanto1103/User-Management-System/tcpserver/responseHandler"
	"log"
)

type UpdateProfileHandler struct{}

func (handler *UpdateProfileHandler) HandleRequest(mapper map[string]string) response.Response {
	profile := mapper[user.CodeProfile]
	userIDFromCookie := mapper[user.CodeCookie]
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
