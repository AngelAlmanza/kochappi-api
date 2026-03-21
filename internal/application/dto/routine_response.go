package dto

import "time"

type RoutineDetailResponse struct {
	ID           int   `json:"id" example:"1"`
	DayOfWeek    int16 `json:"dayOfWeek" example:"1"`
	ExerciseID   int   `json:"exerciseId" example:"1"`
	Sets         int16 `json:"sets" example:"3"`
	Reps         int16 `json:"reps" example:"10"`
	DisplayOrder int16 `json:"displayOrder" example:"1"`
}

type RoutineResponse struct {
	ID         int    `json:"id" example:"1"`
	CustomerID int    `json:"customerId" example:"1"`
	TemplateID *int   `json:"templateId" example:"1"`
	Name       string `json:"name" example:"My Routine"`
	IsActive   bool   `json:"isActive" example:"false"`
}

type RoutineWithDetailsResponse struct {
	ID         int                     `json:"id" example:"1"`
	CustomerID int                     `json:"customerId" example:"1"`
	TemplateID *int                    `json:"templateId" example:"1"`
	Name       string                  `json:"name" example:"My Routine"`
	IsActive   bool                    `json:"isActive" example:"false"`
	Details    []RoutineDetailResponse `json:"details"`
}

type RoutinePeriodResponse struct {
	ID        int        `json:"id" example:"1"`
	RoutineID int        `json:"routineId" example:"1"`
	StartedAt time.Time  `json:"startedAt" example:"2026-01-01T00:00:00Z"`
	EndedAt   *time.Time `json:"endedAt" example:"2026-02-01T00:00:00Z"`
}
