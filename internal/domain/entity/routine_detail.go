package entity

import "time"

type RoutineDetail struct {
	ID           int
	RoutineID    int
	DayOfWeek    int16
	ExerciseID   int
	Sets         int16
	Reps         int16
	DisplayOrder int16
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewRoutineDetail(routineID int, dayOfWeek int16, exerciseID int, sets, reps, displayOrder int16) *RoutineDetail {
	now := time.Now()
	return &RoutineDetail{
		RoutineID:    routineID,
		DayOfWeek:    dayOfWeek,
		ExerciseID:   exerciseID,
		Sets:         sets,
		Reps:         reps,
		DisplayOrder: displayOrder,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
