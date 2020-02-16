package requestHandler

import (
	"github.com/dickynovanto1103/User-Management-System/internal/model"
)

type RequestHandler interface {
	HandleRequest(map[string]string) model.Response
}
