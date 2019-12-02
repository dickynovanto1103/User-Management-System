package response

type Response struct {
	ResponseID int
	Data       map[string]interface{}
}

const (
	ResponseOK           = 1
	ResponseForbidden    = 2
	ResponseCode         = "response"
	ResponseKeyForbidden = "Forbidden"
)
