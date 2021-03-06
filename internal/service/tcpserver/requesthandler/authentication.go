package requesthandler

import (
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"

	"github.com/dickynovanto1103/User-Management-System/internal/common/stringutil"
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/dickynovanto1103/User-Management-System/internal/service/authentication"
)

// AuthenticationHandler authentication request handler
type AuthenticationHandler struct{}

// HandleRequest method for handling authentication request
func (handler *AuthenticationHandler) HandleRequest(mapper map[string]string, redis redis.Redis, db dbsql.DB) model.Response {
	username := mapper[model.CodeUsername]
	password := mapper[model.CodePassword]
	err := authentication.Authenticate(username, password, db)
	mapperResp := make(map[string]string)
	if err != nil {
		if err.Error() == authentication.ErrorNotAuthenticated {
			mapperResp[model.ResponseCode] = authentication.ErrorNotAuthenticated
		}

		mapperResp[model.ResponseCode] = err.Error()
		return model.Response{ResponseID: model.ResponseOK, Data: mapperResp}
	}

	sessionID := stringutil.CreateRandomString(32)
	redis.Set(sessionID, username, 5*time.Hour)
	mapperResp[model.ResponseCode] = sessionID

	return model.Response{ResponseID: model.ResponseOK, Data: mapperResp}
}
