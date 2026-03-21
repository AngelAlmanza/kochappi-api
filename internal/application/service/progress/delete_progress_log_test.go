package progress

import (
	"context"
	"io"
	"testing"
	"time"

	"kochappi/internal/adapter/persistence/mock"
	"kochappi/internal/domain/entity"
)

func TestDeleteProgressLogUseCase_ShouldCascadeDelete(t *testing.T) {
	deletedFiles := []string{}
	deletedPhotos := false
	deletedLog := false

	customerRepo := &mock.MockCustomerRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.Customer, error) {
			return &entity.Customer{ID: id}, nil
		},
	}
	logRepo := &mock.MockLogCustomerProgressRepository{
		GetByIDFn: func(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
			return &entity.LogCustomerProgress{ID: id, CustomerID: 1, CheckDate: time.Now(), Weight: 75}, nil
		},
		DeleteFn: func(ctx context.Context, id int) error {
			deletedLog = true
			return nil
		},
	}
	photoRepo := &mock.MockProgressPhotoRepository{
		GetByLogIDFn: func(ctx context.Context, logID int) ([]entity.ProgressPhoto, error) {
			return []entity.ProgressPhoto{
				{ID: 1, URL: "/uploads/photo1.jpg", LogCustomerProgressID: logID},
				{ID: 2, URL: "/uploads/photo2.jpg", LogCustomerProgressID: logID},
			}, nil
		},
		DeleteByLogIDFn: func(ctx context.Context, logID int) error {
			deletedPhotos = true
			return nil
		},
	}
	fileStorage := &mock.MockFileStorage{
		DeleteFn: func(ctx context.Context, url string) error {
			deletedFiles = append(deletedFiles, url)
			return nil
		},
		UploadFn: func(ctx context.Context, filename string, file io.Reader) (string, error) {
			return "", nil
		},
	}

	uc := NewDeleteProgressLogUseCase(customerRepo, logRepo, photoRepo, fileStorage)
	err := uc.Execute(context.Background(), 1, 1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(deletedFiles) != 2 {
		t.Errorf("Expected 2 files deleted, got %d", len(deletedFiles))
	}
	if !deletedPhotos {
		t.Error("Expected photos to be deleted from DB")
	}
	if !deletedLog {
		t.Error("Expected log to be deleted from DB")
	}
}
