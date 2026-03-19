package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=8,max=72" example:"secureP@ss123"`
	Role     string `json:"role" binding:"required,oneof=trainer client" example:"client" enums:"trainer,client"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secureP@ss123"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email" example:"john@example.com"`
	OTPCode     string `json:"otp_code" binding:"required,len=6" example:"123456"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=72" example:"newSecureP@ss123"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
}
