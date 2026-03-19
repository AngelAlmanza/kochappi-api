package entity

import "time"

type ExerciseRoutine struct {
	ID         int
	ExerciseID int
	RoutineID  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewExerciseRoutine(exerciseID, routineID int) *ExerciseRoutine {
	now := time.Now()
	return &ExerciseRoutine{
		ExerciseID: exerciseID,
		RoutineID:  routineID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
