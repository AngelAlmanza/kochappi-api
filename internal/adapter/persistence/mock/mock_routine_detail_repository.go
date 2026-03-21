package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockRoutineDetailRepository struct {
	GetByRoutineIDFn func(ctx context.Context, routineID int) ([]entity.RoutineDetail, error)
	GetByIDFn        func(ctx context.Context, id int) (*entity.RoutineDetail, error)
	CreateFn         func(ctx context.Context, detail *entity.RoutineDetail) error
	CreateBulkFn     func(ctx context.Context, details []*entity.RoutineDetail) error
	DeleteByIDFn     func(ctx context.Context, id int) error
}

func (r *MockRoutineDetailRepository) GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutineDetail, error) {
	if r.GetByRoutineIDFn != nil {
		return r.GetByRoutineIDFn(ctx, routineID)
	}
	return nil, nil
}

func (r *MockRoutineDetailRepository) GetByID(ctx context.Context, id int) (*entity.RoutineDetail, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockRoutineDetailRepository) Create(ctx context.Context, detail *entity.RoutineDetail) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, detail)
	}
	return nil
}

func (r *MockRoutineDetailRepository) CreateBulk(ctx context.Context, details []*entity.RoutineDetail) error {
	if r.CreateBulkFn != nil {
		return r.CreateBulkFn(ctx, details)
	}
	return nil
}

func (r *MockRoutineDetailRepository) DeleteByID(ctx context.Context, id int) error {
	if r.DeleteByIDFn != nil {
		return r.DeleteByIDFn(ctx, id)
	}
	return nil
}
