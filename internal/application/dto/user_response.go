package dto

import "time"

type UserDetailResponse struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Role      string    `json:"role" example:"trainer"`
	CreatedAt time.Time `json:"createdAt" example:"2026-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2026-01-01T00:00:00Z"`
}
