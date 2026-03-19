package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockTemplateDetailRepository struct {
	GetByTemplateIDFn func(ctx context.Context, templateID int) ([]entity.TemplateDetail, error)
	GetByIDFn         func(ctx context.Context, id int) (*entity.TemplateDetail, error)
	CreateFn          func(ctx context.Context, detail *entity.TemplateDetail) error
	CreateBulkFn      func(ctx context.Context, details []*entity.TemplateDetail) error
	DeleteByIDFn      func(ctx context.Context, id int) error
}

func (r *MockTemplateDetailRepository) GetByTemplateID(ctx context.Context, templateID int) ([]entity.TemplateDetail, error) {
	if r.GetByTemplateIDFn != nil {
		return r.GetByTemplateIDFn(ctx, templateID)
	}
	return nil, nil
}

func (r *MockTemplateDetailRepository) GetByID(ctx context.Context, id int) (*entity.TemplateDetail, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockTemplateDetailRepository) Create(ctx context.Context, detail *entity.TemplateDetail) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, detail)
	}
	return nil
}

func (r *MockTemplateDetailRepository) CreateBulk(ctx context.Context, details []*entity.TemplateDetail) error {
	if r.CreateBulkFn != nil {
		return r.CreateBulkFn(ctx, details)
	}
	return nil
}

func (r *MockTemplateDetailRepository) DeleteByID(ctx context.Context, id int) error {
	if r.DeleteByIDFn != nil {
		return r.DeleteByIDFn(ctx, id)
	}
	return nil
}
