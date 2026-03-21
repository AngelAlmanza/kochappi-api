package routines

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetRoutinePeriodsUseCase struct {
	routineRepo       port.RoutineRepository
	routinePeriodRepo port.RoutinePeriodRepository
}

func NewGetRoutinePeriodsUseCase(
	routineRepo port.RoutineRepository,
	routinePeriodRepo port.RoutinePeriodRepository,
) *GetRoutinePeriodsUseCase {
	return &GetRoutinePeriodsUseCase{
		routineRepo:       routineRepo,
		routinePeriodRepo: routinePeriodRepo,
	}
}

func (uc *GetRoutinePeriodsUseCase) Execute(ctx context.Context, routineID int) ([]dto.RoutinePeriodResponse, error) {
	if _, err := uc.routineRepo.GetByID(ctx, routineID); err != nil {
		return nil, err
	}

	periods, err := uc.routinePeriodRepo.GetByRoutineID(ctx, routineID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.RoutinePeriodResponse, 0, len(periods))
	for _, p := range periods {
		responses = append(responses, dto.RoutinePeriodResponse{
			ID:        p.ID,
			RoutineID: p.RoutineID,
			StartedAt: p.StartedAt,
			EndedAt:   p.EndedAt,
		})
	}
	return responses, nil
}
