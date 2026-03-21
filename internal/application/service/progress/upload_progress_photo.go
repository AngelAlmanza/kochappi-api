package progress

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"

	"github.com/google/uuid"
)

type UploadProgressPhotoUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
	photoRepo    port.ProgressPhotoRepository
	fileStorage  port.FileStorage
}

func NewUploadProgressPhotoUseCase(
	customerRepo port.CustomerRepository,
	logRepo port.LogCustomerProgressRepository,
	photoRepo port.ProgressPhotoRepository,
	fileStorage port.FileStorage,
) *UploadProgressPhotoUseCase {
	return &UploadProgressPhotoUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
		photoRepo:    photoRepo,
		fileStorage:  fileStorage,
	}
}

func (uc *UploadProgressPhotoUseCase) Execute(ctx context.Context, customerID, logID int, pictureType value_object.PictureType, originalFilename string, file io.Reader) (*dto.ProgressPhotoResponse, error) {
	if _, err := uc.customerRepo.GetByID(ctx, customerID); err != nil {
		return nil, err
	}

	log, err := uc.logRepo.GetByID(ctx, logID)
	if err != nil {
		return nil, err
	}

	if log.CustomerID != customerID {
		return nil, &domainerror.ProgressLogNotFoundError{ID: logID}
	}

	ext := filepath.Ext(originalFilename)
	filename := fmt.Sprintf("progress_%d_%s%s", logID, uuid.New().String(), ext)

	url, err := uc.fileStorage.Upload(ctx, filename, file)
	if err != nil {
		return nil, &domainerror.FileUploadError{Message: err.Error()}
	}

	photoEntity := entity.NewProgressPhoto(url, pictureType, logID)
	if err := uc.photoRepo.Create(ctx, photoEntity); err != nil {
		return nil, err
	}

	return &dto.ProgressPhotoResponse{
		ID:          photoEntity.ID,
		URL:         photoEntity.URL,
		PictureType: photoEntity.PictureType.String(),
		LogID:       photoEntity.LogCustomerProgressID,
	}, nil
}
