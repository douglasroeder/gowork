package models

// Result handles API responses
type Result struct {
	StatusCode int         `json:"status_code"`
	Errors     []string    `json:"errors"`
	Payload    interface{} `json:"payload"`
}

// NewResult returns a new instance of result
func NewResult(statusCode int, payload interface{}, errors []string) Result {
	return Result{
		StatusCode: statusCode,
		Payload:    payload,
		Errors:     errors,
	}
}
