package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/common/stringutil"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/service/authentication"
	"time"
)

type AuthenticationHandler struct {}

func (handler *AuthenticationHandler) HandleRequest(mapper map[string]string) model.Response {
	username := mapper[model.CodeUsername]
	password := mapper[model.CodePassword]
	err := authentication.Authenticate(&username, &password)
	mapperResp := make(map[string]string)
	if err != nil {
		if err.Error() == authentication.ErrorNotAuthenticated {
			mapperResp[model.ResponseCode] = authentication.ErrorNotAuthenticated
		} else {
			mapperResp[model.ResponseCode] = err.Error()
		}
	} else {
		sessionID := stringutil.CreateRandomString(32)
		redisutil.Set(sessionID, username, 5*time.Hour)
		mapperResp[model.ResponseCode] = sessionID
	}
	return model.Response{ResponseID: model.ResponseOK, Data: mapperResp}
}
