package http

// W is to wrap JSON-Data for the response
// e.g. `"resources": ["123", "234"]`
type W map[string]any

type Response struct {

	// Status
	Status string `json:"status"`

	// Message is for messages to the api consumer
	Message string `json:"message,omitempty"`

	// Data should used to wrap data which requested the api consumer
	Data W `json:"data,omitempty"`

	// Error is for a specific error message
	Error string `json:"error,omitempty"`

	// Errors is for more specific error message due a specific field
	Errors map[string]string `json:"errors,omitempty"`
}
