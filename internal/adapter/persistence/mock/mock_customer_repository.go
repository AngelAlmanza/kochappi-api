package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockCustomerRepository struct {
	GetAllFn      func(ctx context.Context) ([]entity.Customer, error)
	GetByIDFn     func(ctx context.Context, id int) (*entity.Customer, error)
	GetByUserIDFn func(ctx context.Context, userID int) (*entity.Customer, error)
	CreateFn      func(ctx context.Context, customer *entity.Customer) error
	UpdateFn      func(ctx context.Context, customer *entity.Customer) error
	DeleteFn      func(ctx context.Context, id int) error
}

func (r *MockCustomerRepository) GetAll(ctx context.Context) ([]entity.Customer, error) {
	if r.GetAllFn != nil {
		return r.GetAllFn(ctx)
	}
	return nil, nil
}

func (r *MockCustomerRepository) GetByID(ctx context.Context, id int) (*entity.Customer, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockCustomerRepository) GetByUserID(ctx context.Context, userID int) (*entity.Customer, error) {
	if r.GetByUserIDFn != nil {
		return r.GetByUserIDFn(ctx, userID)
	}
	return nil, nil
}

func (r *MockCustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, customer)
	}
	return nil
}

func (r *MockCustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, customer)
	}
	return nil
}

func (r *MockCustomerRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}
