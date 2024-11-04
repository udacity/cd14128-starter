// utils/response.go
package utils

// MessageResponse defines the structure for success messages
type MessageResponse struct {
    Message string `json:"message" example:"Media deleted"`
}

// HTTPError defines the structure for error responses
type HTTPError struct {
    Code    int    `json:"code" example:"400"`
    Message string `json:"message" example:"Invalid input"`
}
