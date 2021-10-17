package rest

// H is any REST JSON response or payload.
type H map[string]interface{}

// Pagination ...
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// ErrorResponse defines the body of an error response.
type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

// ResponseWithPayload common response with any data in Payload field.
type ResponseWithPayload struct {
	Message    string      `json:"message,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}
