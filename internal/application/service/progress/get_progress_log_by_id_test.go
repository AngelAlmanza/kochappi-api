package progress

import (
	"context"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

func TestGetProgressLogByIDUseCase_ShouldReturnLogWithPhotos(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CheckDate: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), Weight: 75, CustomerID: 1}, nil
		},
	}
	photoRepo := &mock.MockProgressPhotoRepository{
		GetByLogIDFn: func(ctx context.Context, logID int) ([]entity.ProgressPhoto, error) {
			return []entity.ProgressPhoto{
				{ID: 1, URL: "/uploads/photo1.jpg", PictureType: "front", LogCustomerProgressID: logID},
			}, nil
		},
	}

	uc := NewGetProgressLogByIDUseCase(customerRepo, logRepo, photoRepo)
	result, err := uc.Execute(context.Background(), 1, 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if len(result.Photos) != 1 {
		t.Fatalf("Expected 1 photo, got %d", len(result.Photos))
	}
	if result.Photos[0].PictureType != "front" {
		t.Errorf("Expected pictureType front, got %s", result.Photos[0].PictureType)
	}
}

func TestGetProgressLogByIDUseCase_ShouldFailWhenLogNotFound(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return nil, &domainerror.ProgressLogNotFoundError{ID: id}
		},
	}

	uc := NewGetProgressLogByIDUseCase(customerRepo, logRepo, &mock.MockProgressPhotoRepository{})
	_, err := uc.Execute(context.Background(), 1, 99)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressLogNotFoundError); !ok {
		t.Errorf("Expected ProgressLogNotFoundError, got %T", err)
	}
}

func TestGetProgressLogByIDUseCase_ShouldFailWhenLogBelongsToDifferentCustomer(t *testing.T) {
	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CustomerID: 999}, nil
		},
	}

	uc := NewGetProgressLogByIDUseCase(customerRepo, logRepo, &mock.MockProgressPhotoRepository{})
	_, err := uc.Execute(context.Background(), 1, 1)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressLogNotFoundError); !ok {
		t.Errorf("Expected ProgressLogNotFoundError, got %T", err)
	}
}
