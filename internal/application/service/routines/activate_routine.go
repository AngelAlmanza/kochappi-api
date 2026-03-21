package routines

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

type ActivateRoutineUseCase struct {
	routineRepo       port.RoutineRepository
	routinePeriodRepo port.RoutinePeriodRepository
}

func NewActivateRoutineUseCase(
	routineRepo port.RoutineRepository,
	routinePeriodRepo port.RoutinePeriodRepository,
) *ActivateRoutineUseCase {
	return &ActivateRoutineUseCase{
		routineRepo:       routineRepo,
		routinePeriodRepo: routinePeriodRepo,
	}
}

func (uc *ActivateRoutineUseCase) Execute(ctx context.Context, id int) (*dto.RoutineResponse, error) {
	routine, err := uc.routineRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if routine.IsActive {
		return &dto.RoutineResponse{
			ID:         routine.ID,
			CustomerID: routine.CustomerID,
			TemplateID: routine.TemplateID,
			Name:       routine.Name,
			IsActive:   routine.IsActive,
		}, nil
	}

	active, err := uc.routineRepo.GetActiveByCustomerID(ctx, routine.CustomerID)
	if err != nil {
		return nil, err
	}
	// Only one active routine per customer
	if active != nil {
		return nil, &domainerror.ActiveRoutineExistsError{CustomerID: routine.CustomerID}
	}

	routine.Activate()

	if err := uc.routineRepo.Update(ctx, routine); err != nil {
		return nil, err
	}

	// Create a new routine period
	period := entity.NewRoutinePeriod(routine.ID, time.Now())
	if err := uc.routinePeriodRepo.Create(ctx, period); err != nil {
		return nil, err
	}

	return &dto.RoutineResponse{
		ID:         routine.ID,
		CustomerID: routine.CustomerID,
		TemplateID: routine.TemplateID,
		Name:       routine.Name,
		IsActive:   routine.IsActive,
	}, nil
}
