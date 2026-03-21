package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockLogCustomerProgressRepository struct {
	GetByCustomerIDFn func(ctx context.Context, customerID int) ([]entity.LogCustomerProgress, error)
	GetByIDFn         func(ctx context.Context, id int) (*entity.LogCustomerProgress, error)
	CreateFn          func(ctx context.Context, log *entity.LogCustomerProgress) error
	DeleteFn          func(ctx context.Context, id int) error
}

func (r *MockLogCustomerProgressRepository) GetByCustomerID(ctx context.Context, customerID int) ([]entity.LogCustomerProgress, error) {
	if r.GetByCustomerIDFn != nil {
		return r.GetByCustomerIDFn(ctx, customerID)
	}
	return nil, nil
}

func (r *MockLogCustomerProgressRepository) GetByID(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockLogCustomerProgressRepository) Create(ctx context.Context, log *entity.LogCustomerProgress) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, log)
	}
	return nil
}

func (r *MockLogCustomerProgressRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}
