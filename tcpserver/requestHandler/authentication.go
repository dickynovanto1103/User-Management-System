package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/authentication"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/response"
	"github.com/dickynovanto1103/User-Management-System/internal/stringutil"
	"github.com/dickynovanto1103/User-Management-System/internal/user"
	"time"
)

type AuthenticationHandler struct {}

func (handler *AuthenticationHandler) HandleRequest(mapper map[string]string) response.Response {
	username := mapper[user.CodeUsername]
	password := mapper[user.CodePassword]
	err := authentication.Authenticate(&username, &password)
	mapperResp := make(map[string]string)
	if err != nil {
		if err.Error() == authentication.ErrorNotAuthenticated {
			mapperResp[response.ResponseCode] = authentication.ErrorNotAuthenticated
		} else {
			mapperResp[response.ResponseCode] = err.Error()
		}
	} else {
		sessionID := stringutil.CreateRandomString(32)
		redisutil.Set(sessionID, username, 5*time.Hour)
		mapperResp[response.ResponseCode] = sessionID
	}
	return response.Response{ResponseID: response.ResponseOK, Data: mapperResp}
}
