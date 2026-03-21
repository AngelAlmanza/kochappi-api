package progress

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetProgressLogsUseCase struct {
	customerRepo port.CustomerRepository
	logRepo      port.LogCustomerProgressRepository
}

func NewGetProgressLogsUseCase(customerRepo port.CustomerRepository, logRepo port.LogCustomerProgressRepository) *GetProgressLogsUseCase {
	return &GetProgressLogsUseCase{
		customerRepo: customerRepo,
		logRepo:      logRepo,
	}
}

func (uc *GetProgressLogsUseCase) Execute(ctx context.Context, customerID int) ([]dto.ProgressLogResponse, error) {
	if _, err := uc.customerRepo.GetByID(ctx, customerID); err != nil {
		return nil, err
	}

	logs, err := uc.logRepo.GetByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ProgressLogResponse, 0, len(logs))
	for _, l := range logs {
		responses = append(responses, dto.ProgressLogResponse{
			ID:         l.ID,
			CheckDate:  l.CheckDate.Format(time.DateOnly),
			Weight:     l.Weight,
			CustomerID: l.CustomerID,
		})
	}

	return responses, nil
}
