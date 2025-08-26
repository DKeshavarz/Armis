package client

// PutValueRequest is the request body for storing a value
type PutValueRequest struct {
	Value string `json:"value" example:"value"`
}

// ValueResponse is the response for retrieving a key-value
type ValueResponse struct {
	Key   string `json:"key" example:"key"`
	Value string `json:"value" example:"value"`
}

// MessageResponse is a generic success response
type MessageResponse struct {
	Message string `json:"message" example:"key stored"`
	Key     string `json:"key" example:"your-key"`
	Value   string `json:"value,omitempty" example:"your-value"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}