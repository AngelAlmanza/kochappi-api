package entity

import (
	"slices"
	"time"

	domainerror "kochappi/internal/domain/error"
)

type WorkoutStatus string

const (
	WorkoutStatusPending    WorkoutStatus = "pending"
	WorkoutStatusInProgress WorkoutStatus = "in_progress"
	WorkoutStatusCompleted  WorkoutStatus = "completed"
	WorkoutStatusSkipped    WorkoutStatus = "skipped"
)

var validTransitions = map[WorkoutStatus][]WorkoutStatus{
	WorkoutStatusPending:    {WorkoutStatusInProgress, WorkoutStatusSkipped},
	WorkoutStatusInProgress: {WorkoutStatusCompleted, WorkoutStatusSkipped},
}

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

func (ws *WorkoutSession) TransitionTo(next WorkoutStatus) error {
	allowed, ok := validTransitions[ws.Status]
	if !ok {
		return &domainerror.InvalidSessionStatusTransitionError{From: string(ws.Status), To: string(next)}
	}
	if slices.Contains(allowed, next) {
		ws.applyTransition(next)
		return nil
	}
	return &domainerror.InvalidSessionStatusTransitionError{From: string(ws.Status), To: string(next)}
}

func (ws *WorkoutSession) applyTransition(next WorkoutStatus) {
	now := time.Now()
	switch next {
	case WorkoutStatusInProgress:
		ws.StartedAt = &now
	case WorkoutStatusCompleted:
		ws.FinishedAt = &now
	}
	ws.Status = next
	ws.UpdatedAt = now
}
