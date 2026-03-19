package dto

type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	RefreshToken string       `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

type UserResponse struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
	Role  string `json:"role" example:"client"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

type MessageResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

// ErrorResponse represents an API error response.
type ErrorResponse struct {
	Error string `json:"error" example:"descriptive error message"`
	Code  string `json:"code" example:"ERROR_CODE"`
}

// MeResponse represents the response for the authenticated user info endpoint.
type MeResponse struct {
	UserID int    `json:"user_id" example:"1"`
	Role   string `json:"role" example:"client"`
}
