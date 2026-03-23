package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockLogExerciseSessionRepository struct {
	GetByIDFn               func(ctx context.Context, id int) (*entity.LogExerciseSession, error)
	GetByWorkoutSessionIDFn func(ctx context.Context, workoutSessionID int) ([]entity.LogExerciseSession, error)
	CreateFn                func(ctx context.Context, log *entity.LogExerciseSession) error
	UpdateFn                func(ctx context.Context, log *entity.LogExerciseSession) error
	DeleteFn                func(ctx context.Context, id int) error
}

func (r *MockLogExerciseSessionRepository) GetByID(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockLogExerciseSessionRepository) GetByWorkoutSessionID(ctx context.Context, workoutSessionID int) ([]entity.LogExerciseSession, error) {
	if r.GetByWorkoutSessionIDFn != nil {
		return r.GetByWorkoutSessionIDFn(ctx, workoutSessionID)
	}
	return nil, nil
}

func (r *MockLogExerciseSessionRepository) Create(ctx context.Context, log *entity.LogExerciseSession) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, log)
	}
	return nil
}

func (r *MockLogExerciseSessionRepository) Update(ctx context.Context, log *entity.LogExerciseSession) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, log)
	}
	return nil
}

func (r *MockLogExerciseSessionRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}
