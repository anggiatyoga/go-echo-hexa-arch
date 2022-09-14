package webapi

const (
	Success = "success"
	Fail    = "fail"
	Error   = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}
