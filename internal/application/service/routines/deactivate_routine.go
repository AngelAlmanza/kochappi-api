package routines

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type DeactivateRoutineUseCase struct {
	routineRepo       port.RoutineRepository
	routinePeriodRepo port.RoutinePeriodRepository
}

func NewDeactivateRoutineUseCase(
	routineRepo port.RoutineRepository,
	routinePeriodRepo port.RoutinePeriodRepository,
) *DeactivateRoutineUseCase {
	return &DeactivateRoutineUseCase{
		routineRepo:       routineRepo,
		routinePeriodRepo: routinePeriodRepo,
	}
}

func (uc *DeactivateRoutineUseCase) Execute(ctx context.Context, id int) (*dto.RoutineResponse, error) {
	routine, err := uc.routineRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !routine.IsActive {
		return &dto.RoutineResponse{
			ID:         routine.ID,
			CustomerID: routine.CustomerID,
			TemplateID: routine.TemplateID,
			Name:       routine.Name,
			IsActive:   routine.IsActive,
		}, nil
	}

	routine.Deactivate()

	if err := uc.routineRepo.Update(ctx, routine); err != nil {
		return nil, err
	}

	ongoingPeriod, err := uc.routinePeriodRepo.GetOngoingByRoutineID(ctx, routine.ID)
	if err != nil {
		return nil, err
	}
	if ongoingPeriod != nil {
		ongoingPeriod.End(time.Now())
		if err := uc.routinePeriodRepo.Update(ctx, ongoingPeriod); err != nil {
			return nil, err
		}
	}

	return &dto.RoutineResponse{
		ID:         routine.ID,
		CustomerID: routine.CustomerID,
		TemplateID: routine.TemplateID,
		Name:       routine.Name,
		IsActive:   routine.IsActive,
	}, nil
}
