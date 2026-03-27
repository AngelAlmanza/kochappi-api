package dto

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=trainer client"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
