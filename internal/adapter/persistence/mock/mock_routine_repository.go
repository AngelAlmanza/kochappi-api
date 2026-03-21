package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockRoutineRepository struct {
	GetAllFn                func(ctx context.Context) ([]entity.Routine, error)
	GetByIDFn              func(ctx context.Context, id int) (*entity.Routine, error)
	GetByCustomerIDFn      func(ctx context.Context, customerID int) ([]entity.Routine, error)
	GetActiveByCustomerIDFn func(ctx context.Context, customerID int) (*entity.Routine, error)
	CreateFn               func(ctx context.Context, routine *entity.Routine) error
	UpdateFn               func(ctx context.Context, routine *entity.Routine) error
	GetAllActiveFn         func(ctx context.Context) ([]entity.Routine, error)
}

func (r *MockRoutineRepository) GetAll(ctx context.Context) ([]entity.Routine, error) {
	if r.GetAllFn != nil {
		return r.GetAllFn(ctx)
	}
	return nil, nil
}

func (r *MockRoutineRepository) GetByID(ctx context.Context, id int) (*entity.Routine, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockRoutineRepository) GetByCustomerID(ctx context.Context, customerID int) ([]entity.Routine, error) {
	if r.GetByCustomerIDFn != nil {
		return r.GetByCustomerIDFn(ctx, customerID)
	}
	return nil, nil
}

func (r *MockRoutineRepository) GetActiveByCustomerID(ctx context.Context, customerID int) (*entity.Routine, error) {
	if r.GetActiveByCustomerIDFn != nil {
		return r.GetActiveByCustomerIDFn(ctx, customerID)
	}
	return nil, nil
}

func (r *MockRoutineRepository) Create(ctx context.Context, routine *entity.Routine) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, routine)
	}
	return nil
}

func (r *MockRoutineRepository) Update(ctx context.Context, routine *entity.Routine) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, routine)
	}
	return nil
}

func (r *MockRoutineRepository) GetAllActive(ctx context.Context) ([]entity.Routine, error) {
	if r.GetAllActiveFn != nil {
		return r.GetAllActiveFn(ctx)
	}
	return nil, nil
}
