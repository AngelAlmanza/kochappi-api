package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockTemplateRepository struct {
	GetAllFn  func(ctx context.Context) ([]entity.Template, error)
	GetByIDFn func(ctx context.Context, id int) (*entity.Template, error)
	CreateFn  func(ctx context.Context, template *entity.Template) error
	UpdateFn  func(ctx context.Context, template *entity.Template) error
	DeleteFn  func(ctx context.Context, id int) error
}

func (r *MockTemplateRepository) GetAll(ctx context.Context) ([]entity.Template, error) {
	if r.GetAllFn != nil {
		return r.GetAllFn(ctx)
	}
	return nil, nil
}

func (r *MockTemplateRepository) GetByID(ctx context.Context, id int) (*entity.Template, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockTemplateRepository) Create(ctx context.Context, template *entity.Template) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, template)
	}
	return nil
}

func (r *MockTemplateRepository) Update(ctx context.Context, template *entity.Template) error {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, template)
	}
	return nil
}

func (r *MockTemplateRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}
