package mock

import (
	"context"
	"kochappi/internal/domain/entity"
)

type MockUserRepository struct {
	GetAllFn     func(ctx context.Context, role *entity.Role) ([]entity.User, error)
	CreateFn     func(ctx context.Context, user *entity.User) error
	GetByIDFn    func(ctx context.Context, id int) (*entity.User, error)
	GetByEmailFn func(ctx context.Context, email string) (*entity.User, error)
	UpdateFn     func(ctx context.Context, user *entity.User) error
}

func (r *MockUserRepository) GetAll(ctx context.Context, role *entity.Role) ([]entity.User, error) {
	if r.GetAllFn != nil {
		return r.GetAllFn(ctx, role)
	}
	return nil, nil
}

func (r *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, user)
	}
	return nil
}

func (r *MockUserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	if r.GetByEmailFn != nil {
		return r.GetByEmailFn(ctx, email)
	}
	return nil, nil
}

func (r *MockUserRepository) Update(ctx context.Context, user *entity.User) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, user)
	}
	return nil
}
