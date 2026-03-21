package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockRoutinePeriodRepository struct {
	GetByRoutineIDFn        func(ctx context.Context, routineID int) ([]entity.RoutinePeriod, error)
	GetOngoingByRoutineIDFn func(ctx context.Context, routineID int) (*entity.RoutinePeriod, error)
	CreateFn                func(ctx context.Context, period *entity.RoutinePeriod) error
	UpdateFn                func(ctx context.Context, period *entity.RoutinePeriod) error
}

func (r *MockRoutinePeriodRepository) GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutinePeriod, error) {
	if r.GetByRoutineIDFn != nil {
		return r.GetByRoutineIDFn(ctx, routineID)
	}
	return nil, nil
}

func (r *MockRoutinePeriodRepository) GetOngoingByRoutineID(ctx context.Context, routineID int) (*entity.RoutinePeriod, error) {
	if r.GetOngoingByRoutineIDFn != nil {
		return r.GetOngoingByRoutineIDFn(ctx, routineID)
	}
	return nil, nil
}

func (r *MockRoutinePeriodRepository) Create(ctx context.Context, period *entity.RoutinePeriod) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, period)
	}
	return nil
}

func (r *MockRoutinePeriodRepository) Update(ctx context.Context, period *entity.RoutinePeriod) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, period)
	}
	return nil
}
