package mock

import (
	"context"

	"kochappi/internal/domain/entity"
)

type MockProgressPhotoRepository struct {
	GetByLogIDFn    func(ctx context.Context, logID int) ([]entity.ProgressPhoto, error)
	GetByIDFn       func(ctx context.Context, id int) (*entity.ProgressPhoto, error)
	CreateFn        func(ctx context.Context, photo *entity.ProgressPhoto) error
	DeleteFn        func(ctx context.Context, id int) error
	DeleteByLogIDFn func(ctx context.Context, logID int) error
}

func (r *MockProgressPhotoRepository) GetByLogID(ctx context.Context, logID int) ([]entity.ProgressPhoto, error) {
	if r.GetByLogIDFn != nil {
		return r.GetByLogIDFn(ctx, logID)
	}
	return nil, nil
}

func (r *MockProgressPhotoRepository) GetByID(ctx context.Context, id int) (*entity.ProgressPhoto, error) {
	if r.GetByIDFn != nil {
		return r.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (r *MockProgressPhotoRepository) Create(ctx context.Context, photo *entity.ProgressPhoto) error {
	if r.CreateFn != nil {
		return r.CreateFn(ctx, photo)
	}
	return nil
}

func (r *MockProgressPhotoRepository) Delete(ctx context.Context, id int) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}

func (r *MockProgressPhotoRepository) DeleteByLogID(ctx context.Context, logID int) error {
	if r.DeleteByLogIDFn != nil {
		return r.DeleteByLogIDFn(ctx, logID)
	}
	return nil
}
