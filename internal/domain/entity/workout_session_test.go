package entity

import (
	"testing"
	"time"

	domainerror "kochappi/internal/domain/error"
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

func TestWorkoutSession_TransitionTo_PendingToInProgress(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())

	if err := session.TransitionTo(WorkoutStatusInProgress); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if session.Status != WorkoutStatusInProgress {
		t.Errorf("Expected status %q, got %q", WorkoutStatusInProgress, session.Status)
	}
	if session.StartedAt == nil {
		t.Error("Expected StartedAt to be set after transitioning to in_progress")
	}
}

func TestWorkoutSession_TransitionTo_InProgressToCompleted(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	_ = session.TransitionTo(WorkoutStatusInProgress)

	if err := session.TransitionTo(WorkoutStatusCompleted); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if session.Status != WorkoutStatusCompleted {
		t.Errorf("Expected status %q, got %q", WorkoutStatusCompleted, session.Status)
	}
	if session.FinishedAt == nil {
		t.Error("Expected FinishedAt to be set after transitioning to completed")
	}
}

func TestWorkoutSession_TransitionTo_PendingToSkipped(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())

	if err := session.TransitionTo(WorkoutStatusSkipped); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if session.Status != WorkoutStatusSkipped {
		t.Errorf("Expected status %q, got %q", WorkoutStatusSkipped, session.Status)
	}
	if session.FinishedAt != nil {
		t.Error("Expected FinishedAt to remain nil after skipping")
	}
}

func TestWorkoutSession_TransitionTo_InProgressToSkipped(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	_ = session.TransitionTo(WorkoutStatusInProgress)

	if err := session.TransitionTo(WorkoutStatusSkipped); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if session.Status != WorkoutStatusSkipped {
		t.Errorf("Expected status %q, got %q", WorkoutStatusSkipped, session.Status)
	}
}

func TestWorkoutSession_TransitionTo_InvalidFromCompleted(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	_ = session.TransitionTo(WorkoutStatusInProgress)
	_ = session.TransitionTo(WorkoutStatusCompleted)

	err := session.TransitionTo(WorkoutStatusSkipped)
	if err == nil {
		t.Fatal("Expected error when transitioning from completed, got nil")
	}
	if _, ok := err.(*domainerror.InvalidSessionStatusTransitionError); !ok {
		t.Errorf("Expected InvalidSessionStatusTransitionError, got %T", err)
	}
}

func TestWorkoutSession_TransitionTo_InvalidFromSkipped(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())
	_ = session.TransitionTo(WorkoutStatusSkipped)

	err := session.TransitionTo(WorkoutStatusInProgress)
	if err == nil {
		t.Fatal("Expected error when transitioning from skipped, got nil")
	}
	if _, ok := err.(*domainerror.InvalidSessionStatusTransitionError); !ok {
		t.Errorf("Expected InvalidSessionStatusTransitionError, got %T", err)
	}
}

func TestWorkoutSession_TransitionTo_InvalidPendingToCompleted(t *testing.T) {
	session := NewWorkoutSession(1, 1, time.Now())

	err := session.TransitionTo(WorkoutStatusCompleted)
	if err == nil {
		t.Fatal("Expected error when skipping in_progress, got nil")
	}
}
