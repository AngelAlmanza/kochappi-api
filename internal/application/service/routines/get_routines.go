package routines

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type GetRoutinesUseCase struct {
	routineRepo port.RoutineRepository
}

func NewGetRoutinesUseCase(routineRepo port.RoutineRepository) *GetRoutinesUseCase {
	return &GetRoutinesUseCase{routineRepo: routineRepo}
}

func (uc *GetRoutinesUseCase) Execute(ctx context.Context, customerID *int) ([]dto.RoutineResponse, error) {
	var routines []entity.Routine
	var err error

	if customerID != nil {
		routines, err = uc.routineRepo.GetByCustomerID(ctx, *customerID)
	} else {
		routines, err = uc.routineRepo.GetAll(ctx)
	}
	if err != nil {
		return nil, err
	}

	responses := make([]dto.RoutineResponse, 0, len(routines))
	for _, r := range routines {
		responses = append(responses, dto.RoutineResponse{
			ID:         r.ID,
			CustomerID: r.CustomerID,
			TemplateID: r.TemplateID,
			Name:       r.Name,
			IsActive:   r.IsActive,
		})
	}
	return responses, nil
}
