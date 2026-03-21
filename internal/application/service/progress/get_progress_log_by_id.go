package progress

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	domainerror "kochappi/internal/domain/error"
)

type GetProgressLogByIDUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
	photoRepo    port.ProgressPhotoRepository
}

func NewGetProgressLogByIDUseCase(customerRepo port.CustomerRepository, logRepo port.LogCustomerProgressRepository, photoRepo port.ProgressPhotoRepository) *GetProgressLogByIDUseCase {
	return &GetProgressLogByIDUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
		photoRepo:    photoRepo,
	}
}

func (uc *GetProgressLogByIDUseCase) Execute(ctx context.Context, customerID, logID int) (*dto.ProgressLogWithPhotosResponse, error) {
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

	photos, err := uc.photoRepo.GetByLogID(ctx, logID)
	if err != nil {
		return nil, err
	}

	photoResponses := make([]dto.ProgressPhotoResponse, 0, len(photos))
	for _, p := range photos {
		photoResponses = append(photoResponses, dto.ProgressPhotoResponse{
			ID:          p.ID,
			URL:         p.URL,
			PictureType: p.PictureType.String(),
			LogID:       p.LogCustomerProgressID,
		})
	}

	return &dto.ProgressLogWithPhotosResponse{
		ID:         log.ID,
		CheckDate:  log.CheckDate.Format(time.DateOnly),
		Weight:     log.Weight,
		CustomerID: log.CustomerID,
		Photos:     photoResponses,
	}, nil
}
