package progress

import (
	"context"

	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type DeleteProgressPhotoUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
	photoRepo    port.ProgressPhotoRepository
	fileStorage  port.FileStorage
}

func NewDeleteProgressPhotoUseCase(
	customerRepo port.CustomerRepository,
	logRepo port.LogCustomerProgressRepository,
	photoRepo port.ProgressPhotoRepository,
	fileStorage port.FileStorage,
) *DeleteProgressPhotoUseCase {
	return &DeleteProgressPhotoUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
		photoRepo:    photoRepo,
		fileStorage:  fileStorage,
	}
}

func (uc *DeleteProgressPhotoUseCase) Execute(ctx context.Context, customerID, logID, photoID int) error {
	if _, err := uc.customerRepo.GetByID(ctx, customerID); err != nil {
		return err
	}

	log, err := uc.logRepo.GetByID(ctx, logID)
	if err != nil {
		return err
	}

	if log.CustomerID != customerID {
		return &domainerror.ProgressLogNotFoundError{ID: logID}
	}

	photo, err := uc.photoRepo.GetByID(ctx, photoID)
	if err != nil {
		return err
	}

	if photo.LogCustomerProgressID != logID {
		return &domainerror.ProgressPhotoNotFoundError{ID: photoID}
	}

	if err := uc.photoRepo.Delete(ctx, photoID); err != nil {
		return err
	}

	if err := uc.fileStorage.Delete(ctx, photo.URL); err != nil {
		return err
	}

	return nil
}
