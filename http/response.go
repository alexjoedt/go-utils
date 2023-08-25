package http

// Wrap is to wrap JSON-Data for the response
// e.g. `"resources": ["123", "234"]`
type Wrap map[string]any

type Response struct {
	Status  ResponseStatus `json:"status"`
	Code    string         `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    Wrap           `json:"data,omitempty"`
	Errors  any            `json:"errors,omitempty"`
}

type ResponseStatus string

const (
	StatusOK    ResponseStatus = "ok"
	StatusWarn  ResponseStatus = "warning"
	StatusErr   ResponseStatus = "error"
	StatusFatal ResponseStatus = "fatal"
)
