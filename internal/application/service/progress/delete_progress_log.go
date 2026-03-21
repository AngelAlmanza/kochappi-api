package progress

import (
	"context"

	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type DeleteProgressLogUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
	photoRepo    port.ProgressPhotoRepository
	fileStorage  port.FileStorage
}

func NewDeleteProgressLogUseCase(
	customerRepo port.CustomerRepository,
	logRepo port.LogCustomerProgressRepository,
	photoRepo port.ProgressPhotoRepository,
	fileStorage port.FileStorage,
) *DeleteProgressLogUseCase {
	return &DeleteProgressLogUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
		photoRepo:    photoRepo,
		fileStorage:  fileStorage,
	}
}

func (uc *DeleteProgressLogUseCase) Execute(ctx context.Context, customerID, logID int) error {
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

	photos, err := uc.photoRepo.GetByLogID(ctx, logID)
	if err != nil {
		return err
	}

	for _, photo := range photos {
		_ = uc.fileStorage.Delete(ctx, photo.URL)
	}

	if err := uc.photoRepo.DeleteByLogID(ctx, logID); err != nil {
		return err
	}

	return uc.logRepo.Delete(ctx, logID)
}
