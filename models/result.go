package models

// Result handles API responses
type Result struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Payload    interface{} `json:"payload"`
}

// NewResult returns a new instance of result
func NewResult(statusCode int, payload interface{}, error string) Result {
	return Result{
		StatusCode: statusCode,
		Error:      error,
		Payload:    payload,
	}
}
