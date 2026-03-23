package dto

type UpdateWorkoutSessionStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=in_progress completed skipped" example:"in_progress"`
}

type CreateExerciseLogRequest struct {
	RoutineDetailID int     `json:"routineDetailId" binding:"required" example:"1"`
	SetNumber       int16   `json:"setNumber" binding:"required,min=1" example:"1"`
	RepsDone        int16   `json:"repsDone" binding:"required,min=0" example:"10"`
	Weight          float64 `json:"weight" binding:"required,min=0" example:"50.0"`
	Notes           string  `json:"notes" example:"Felt strong"`
}

type UpdateExerciseLogRequest struct {
	SetNumber int16   `json:"setNumber" binding:"required,min=1" example:"1"`
	RepsDone  int16   `json:"repsDone" binding:"required,min=0" example:"10"`
	Weight    float64 `json:"weight" binding:"required,min=0" example:"50.0"`
	Notes     string  `json:"notes" example:"Felt strong"`
}
