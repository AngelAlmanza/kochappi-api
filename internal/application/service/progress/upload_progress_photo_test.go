package progress

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"
)

func TestUploadProgressPhotoUseCase_ShouldUploadPhoto(t *testing.T) {
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
		CreateFn: func(ctx context.Context, photo *entity.ProgressPhoto) error {
			photo.ID = 1
			return nil
		},
	}
	fileStorage := &mock.MockFileStorage{
		UploadFn: func(ctx context.Context, filename string, file io.Reader) (string, error) {
			return "/uploads/" + filename, nil
		},
	}

	uc := NewUploadProgressPhotoUseCase(customerRepo, logRepo, photoRepo, fileStorage)
	result, err := uc.Execute(context.Background(), 1, 1, value_object.PictureTypeFront, "photo.jpg", strings.NewReader("fake-image-data"))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.PictureType != "front" {
		t.Errorf("Expected pictureType front, got %s", result.PictureType)
	}
}

func TestUploadProgressPhotoUseCase_ShouldFailWhenLogNotFound(t *testing.T) {
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

	uc := NewUploadProgressPhotoUseCase(customerRepo, logRepo, &mock.MockProgressPhotoRepository{}, &mock.MockFileStorage{})
	_, err := uc.Execute(context.Background(), 1, 99, value_object.PictureTypeFront, "photo.jpg", strings.NewReader("data"))

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressLogNotFoundError); !ok {
		t.Errorf("Expected ProgressLogNotFoundError, got %T", err)
	}
}

func TestUploadProgressPhotoUseCase_ShouldFailWhenLogBelongsToDifferentCustomer(t *testing.T) {
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

	uc := NewUploadProgressPhotoUseCase(customerRepo, logRepo, &mock.MockProgressPhotoRepository{}, &mock.MockFileStorage{})
	_, err := uc.Execute(context.Background(), 1, 1, value_object.PictureTypeFront, "photo.jpg", strings.NewReader("data"))

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*domainerror.ProgressLogNotFoundError); !ok {
		t.Errorf("Expected ProgressLogNotFoundError, got %T", err)
	}
}
