package dto

type CreateTemplateDetailRequest struct {
	DayOfWeek    int16 `json:"dayOfWeek" binding:"required,min=0,max=6" example:"1"`
	ExerciseID   int   `json:"exerciseId" binding:"required" example:"1"`
	Sets         int16 `json:"sets" binding:"required" example:"3"`
	Reps         int16 `json:"reps" binding:"required" example:"10"`
	DisplayOrder int16 `json:"displayOrder" binding:"required" example:"1"`
}

type CreateTemplateRequest struct {
	Name        string                        `json:"name" binding:"required" example:"Push/Pull/Legs"`
	Description string                        `json:"description" example:"Classic PPL split"`
	Details     []CreateTemplateDetailRequest `json:"details"`
}

type UpdateTemplateRequest struct {
	Name        string `json:"name" binding:"required" example:"Push/Pull/Legs"`
	Description string `json:"description" example:"Classic PPL split"`
}

type AddTemplateDetailRequest struct {
	DayOfWeek    int16 `json:"dayOfWeek" binding:"required,min=0,max=6" example:"1"`
	ExerciseID   int   `json:"exerciseId" binding:"required" example:"1"`
	Sets         int16 `json:"sets" binding:"required" example:"3"`
	Reps         int16 `json:"reps" binding:"required" example:"10"`
	DisplayOrder int16 `json:"displayOrder" binding:"required" example:"1"`
}
