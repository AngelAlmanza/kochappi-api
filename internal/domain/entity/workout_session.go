package entity

import "time"

type WorkoutStatus string

const (
	WorkoutStatusPending    WorkoutStatus = "pending"
	WorkoutStatusInProgress WorkoutStatus = "in_progress"
	WorkoutStatusCompleted  WorkoutStatus = "completed"
	WorkoutStatusSkipped    WorkoutStatus = "skipped"
)

type WorkoutSession struct {
	ID           int
	RoutineID    int
	ScheduledDay int16
	ActualDate   time.Time
	Status       WorkoutStatus
	StartedAt    *time.Time
	FinishedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewWorkoutSession(routineID int, scheduledDay int16, actualDate time.Time) *WorkoutSession {
	now := time.Now()
	return &WorkoutSession{
		RoutineID:    routineID,
		ScheduledDay: scheduledDay,
		ActualDate:   actualDate,
		Status:       WorkoutStatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (ws *WorkoutSession) Start() {
	now := time.Now()
	ws.Status = WorkoutStatusInProgress
	ws.StartedAt = &now
	ws.UpdatedAt = now
}

func (ws *WorkoutSession) Complete() {
	now := time.Now()
	ws.Status = WorkoutStatusCompleted
	ws.FinishedAt = &now
	ws.UpdatedAt = now
}

func (ws *WorkoutSession) Skip() {
	ws.Status = WorkoutStatusSkipped
	ws.UpdatedAt = time.Now()
}
