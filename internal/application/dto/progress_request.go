package dto

type CreateProgressLogRequest struct {
	CheckDate string `json:"checkDate" binding:"required" example:"2024-01-15"`
	Weight    int    `json:"weight" binding:"required" example:"75"`
}
