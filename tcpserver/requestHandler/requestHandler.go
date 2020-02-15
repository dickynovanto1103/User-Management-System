package requestHandler

import "github.com/dickynovanto1103/User-Management-System/internal/response"

type RequestHandler interface {
	HandleRequest(map[string]string) response.Response
}
