package dto

type TemplateDetailResponse struct {
	ID           int   `json:"id" example:"1"`
	DayOfWeek    int16 `json:"dayOfWeek" example:"1"`
	ExerciseID   int   `json:"exerciseId" example:"1"`
	Sets         int16 `json:"sets" example:"3"`
	Reps         int16 `json:"reps" example:"10"`
	DisplayOrder int16 `json:"displayOrder" example:"1"`
}

type TemplateResponse struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Push/Pull/Legs"`
	Description string `json:"description" example:"Classic PPL split"`
}

type TemplateWithDetailsResponse struct {
	ID          int                      `json:"id" example:"1"`
	Name        string                   `json:"name" example:"Push/Pull/Legs"`
	Description string                   `json:"description" example:"Classic PPL split"`
	Details     []TemplateDetailResponse `json:"details"`
}
