package dto

type CreateRoutineDetailRequest struct {
	DayOfWeek    int16 `json:"dayOfWeek" binding:"required,min=0,max=6" example:"1"`
	ExerciseID   int   `json:"exerciseId" binding:"required" example:"1"`
	Sets         int16 `json:"sets" binding:"required" example:"3"`
	Reps         int16 `json:"reps" binding:"required" example:"10"`
	DisplayOrder int16 `json:"displayOrder" binding:"required" example:"1"`
}

type CreateRoutineRequest struct {
	CustomerID int                          `json:"customerId" binding:"required" example:"1"`
	TemplateID *int                         `json:"templateId" example:"1"`
	Name       string                       `json:"name" binding:"required" example:"My Routine"`
	IsActive   bool                         `json:"isActive" example:"false"`
	Details    []CreateRoutineDetailRequest `json:"details"`
}

type UpdateRoutineRequest struct {
	Name string `json:"name" binding:"required" example:"My Routine"`
}

type AddRoutineDetailRequest struct {
	DayOfWeek    int16 `json:"dayOfWeek" binding:"required,min=0,max=6" example:"1"`
	ExerciseID   int   `json:"exerciseId" binding:"required" example:"1"`
	Sets         int16 `json:"sets" binding:"required" example:"3"`
	Reps         int16 `json:"reps" binding:"required" example:"10"`
	DisplayOrder int16 `json:"displayOrder" binding:"required" example:"1"`
}
