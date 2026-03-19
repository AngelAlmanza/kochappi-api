package entity

import "time"

type LogExerciseSession struct {
	ID               int
	WorkoutSessionID int
	RoutineDetailID  int
	SetNumber        int16
	RepsDone         int16
	Weight           float64
	Notes            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewLogExerciseSession(workoutSessionID, routineDetailID int, setNumber, repsDone int16, weight float64, notes string) *LogExerciseSession {
	now := time.Now()
	return &LogExerciseSession{
		WorkoutSessionID: workoutSessionID,
		RoutineDetailID:  routineDetailID,
		SetNumber:        setNumber,
		RepsDone:         repsDone,
		Weight:           weight,
		Notes:            notes,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
