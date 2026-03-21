package progress

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

type CreateProgressLogUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
}

func NewCreateProgressLogUseCase(customerRepo port.CustomerRepository, logRepo port.LogCustomerProgressRepository) *CreateProgressLogUseCase {
	return &CreateProgressLogUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
	}
}

func (uc *CreateProgressLogUseCase) Execute(ctx context.Context, customerID int, req *dto.CreateProgressLogRequest) (*dto.ProgressLogResponse, error) {
	if _, err := uc.customerRepo.GetByID(ctx, customerID); err != nil {
		return nil, err
	}

	checkDate, err := time.Parse(time.DateOnly, req.CheckDate)
	if err != nil {
		return nil, &domainerror.InvalidCheckDateError{CheckDate: req.CheckDate}
	}

	if req.Weight <= 0 {
		return nil, &domainerror.InvalidWeightError{}
	}

	logEntity := entity.NewLogCustomerProgress(customerID, checkDate, req.Weight)
	if err := uc.logRepo.Create(ctx, logEntity); err != nil {
		return nil, err
	}

	return &dto.ProgressLogResponse{
		ID:         logEntity.ID,
		CheckDate:  logEntity.CheckDate.Format(time.DateOnly),
		Weight:     logEntity.Weight,
		CustomerID: logEntity.CustomerID,
	}, nil
}
