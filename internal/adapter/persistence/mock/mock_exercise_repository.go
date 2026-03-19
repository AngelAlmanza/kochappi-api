package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockExerciseRepository struct {
	GetAllFn  func(ctx context.Context) ([]entity.Exercise, error)
	GetByIDFn func(ctx context.Context, id int) (*entity.Exercise, error)
	CreateFn  func(ctx context.Context, exercise *entity.Exercise) error
	UpdateFn  func(ctx context.Context, exercise *entity.Exercise) error
	DeleteFn  func(ctx context.Context, id int) error
}

func (r *MockExerciseRepository) GetAll(ctx context.Context) ([]entity.Exercise, error) {
	if r.GetAllFn != nil {
		return r.GetAllFn(ctx)
	}
	return nil, nil
}

func (r *MockExerciseRepository) GetByID(ctx context.Context, id int) (*entity.Exercise, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockExerciseRepository) Create(ctx context.Context, exercise *entity.Exercise) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, exercise)
	}
	return nil
}

func (r *MockExerciseRepository) Update(ctx context.Context, exercise *entity.Exercise) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, exercise)
	}
	return nil
}

func (r *MockExerciseRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}
