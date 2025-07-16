package dto

// ErrorResponse is a standard response for general errors.
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// ValidationErrorResponse is a standard response for validation errors.
type ValidationErrorResponse struct {
	Error   string                 `json:"error"`
	Details map[string]interface{} `json:"details"`
}
