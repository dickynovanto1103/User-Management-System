package model

type Response struct {
	ResponseID int32
	Data       map[string]string
}

const (
	ResponseOK           = 1
	ResponseForbidden    = 2
	ResponseError        = 3
	ResponseCode         = "response"
	ResponseKeyForbidden = "forbidden"
	ResponseKeyError     = "error"
)
