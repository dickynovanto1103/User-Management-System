package requesthandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/model"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
)

type RequestHandler interface {
	HandleRequest(map[string]string, redis.Redis, dbsql.DB) model.Response
}
