package entity

import "time"

type TemplateDetail struct {
	ID           int
	TemplateID   int
	DayOfWeek    int16
	ExerciseID   int
	Sets         int16
	Reps         int16
	DisplayOrder int16
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTemplateDetail(templateID int, dayOfWeek int16, exerciseID int, sets, reps, displayOrder int16) *TemplateDetail {
	now := time.Now()
	return &TemplateDetail{
		TemplateID:   templateID,
		DayOfWeek:    dayOfWeek,
		ExerciseID:   exerciseID,
		Sets:         sets,
		Reps:         reps,
		DisplayOrder: displayOrder,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
