package responses

import "crud_operation/dto"

// UserResponse represents the response structure for user-related operations.
type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message,omitempty"`
	Error   error     `json:"error,omitempty"`
	Data    *dto.User `json:"data,omitempty"`
}

// Struct to return all users
type UsersResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Error   error       `json:"error,omitempty"`
	Data    *[]dto.User `json:"data,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
