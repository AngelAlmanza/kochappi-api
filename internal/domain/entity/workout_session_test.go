package entity

import (
	"testing"
	"time"
)

func TestWorkoutSession_NewSession_ShouldHavePendingStatus(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())

	if session.Status != WorkoutStatusPending {
		t.Errorf("Expected status %q, got %q", WorkoutStatusPending, session.Status)
	}
	if session.StartedAt != nil {
		t.Error("Expected StartedAt to be nil on new session")
	}
	if session.FinishedAt != nil {
		t.Error("Expected FinishedAt to be nil on new session")
	}
}

func TestWorkoutSession_Start_ShouldSetStatusAndTimestamp(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	session.Start()

	if session.Status != WorkoutStatusInProgress {
		t.Errorf("Expected status %q, got %q", WorkoutStatusInProgress, session.Status)
	}
	if session.StartedAt == nil {
		t.Error("Expected StartedAt to be set after Start()")
	}
}

func TestWorkoutSession_Complete_ShouldSetStatusAndTimestamp(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	session.Start()
	session.Complete()

	if session.Status != WorkoutStatusCompleted {
		t.Errorf("Expected status %q, got %q", WorkoutStatusCompleted, session.Status)
	}
	if session.FinishedAt == nil {
		t.Error("Expected FinishedAt to be set after Complete()")
	}
}

func TestWorkoutSession_Skip_ShouldSetStatus(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	session.Skip()

	if session.Status != WorkoutStatusSkipped {
		t.Errorf("Expected status %q, got %q", WorkoutStatusSkipped, session.Status)
	}
	if session.FinishedAt != nil {
		t.Error("Expected FinishedAt to remain nil after Skip()")
	}
}
