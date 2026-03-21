package progress

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestDeleteProgressPhotoUseCase_ShouldDeletePhoto(t *testing.T) {
	deletedFileURL := ""
	deletedPhotoID := 0

	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CustomerID: 1, CheckDate: time.Now(), Weight: 75}, nil
		},
	}
	photoRepo := &mock.MockProgressPhotoRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.ProgressPhoto, error) {
			return &entity.ProgressPhoto{ID: id, URL: "/uploads/photo.jpg", PictureType: "front", LogCustomerProgressID: 1}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deletedPhotoID = id
			return nil
		},
	}
	fileStorage := &mock.MockFileStorage{
		DeleteFn: func(ctx context.Context, url string) error {
			deletedFileURL = url
			return nil
		},
	}

	uc := NewDeleteProgressPhotoUseCase(customerRepo, logRepo, photoRepo, fileStorage)
	err := uc.Execute(context.Background(), 1, 1, 5)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if deletedPhotoID != 5 {
		t.Errorf("Expected photo ID 5 deleted, got %d", deletedPhotoID)
	}
	if deletedFileURL != "/uploads/photo.jpg" {
		t.Errorf("Expected file /uploads/photo.jpg deleted, got %s", deletedFileURL)
	}
}

func TestDeleteProgressPhotoUseCase_ShouldFailWhenPhotoNotFound(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CustomerID: 1, CheckDate: time.Now(), Weight: 75}, nil
		},
	}
	photoRepo := &mock.MockProgressPhotoRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.ProgressPhoto, error) {
			return nil, &domainerror.ProgressPhotoNotFoundError{ID: id}
		},
	}

	uc := NewDeleteProgressPhotoUseCase(customerRepo, logRepo, photoRepo, &mock.MockFileStorage{})
	err := uc.Execute(context.Background(), 1, 1, 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressPhotoNotFoundError); !ok {
		t.Errorf("Expected ProgressPhotoNotFoundError, got %T", err)
	}
}

func TestDeleteProgressPhotoUseCase_ShouldFailWhenPhotoBelongsToDifferentLog(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CustomerID: 1, CheckDate: time.Now(), Weight: 75}, nil
		},
	}
	photoRepo := &mock.MockProgressPhotoRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.ProgressPhoto, error) {
			return &entity.ProgressPhoto{ID: id, LogCustomerProgressID: 999}, nil
		},
	}

	uc := NewDeleteProgressPhotoUseCase(customerRepo, logRepo, photoRepo, &mock.MockFileStorage{})
	err := uc.Execute(context.Background(), 1, 1, 5)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressPhotoNotFoundError); !ok {
		t.Errorf("Expected ProgressPhotoNotFoundError, got %T", err)
	}
}
