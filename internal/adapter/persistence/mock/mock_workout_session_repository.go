package mock

import (
	"context"

	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type MockWorkoutSessionRepository struct {
	GetByIDFn       func(ctx context.Context, id int) (*entity.WorkoutSession, error)
	GetByCriteriaFn func(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error)
	CreateFn        func(ctx context.Context, session *entity.WorkoutSession) error
	CreateBulkFn    func(ctx context.Context, sessions []*entity.WorkoutSession) error
	UpdateFn        func(ctx context.Context, session *entity.WorkoutSession) error
}

func (r *MockWorkoutSessionRepository) GetByID(ctx context.Context, id int) (*entity.WorkoutSession, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockWorkoutSessionRepository) GetByCriteria(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error) {
	if r.GetByCriteriaFn != nil {
		return r.GetByCriteriaFn(ctx, criteria)
	}
	return nil, nil
}

func (r *MockWorkoutSessionRepository) Create(ctx context.Context, session *entity.WorkoutSession) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, session)
	}
	return nil
}

func (r *MockWorkoutSessionRepository) CreateBulk(ctx context.Context, sessions []*entity.WorkoutSession) error {
	if r.CreateBulkFn != nil {
		return r.CreateBulkFn(ctx, sessions)
	}
	return nil
}

func (r *MockWorkoutSessionRepository) Update(ctx context.Context, session *entity.WorkoutSession) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, session)
	}
	return nil
}
