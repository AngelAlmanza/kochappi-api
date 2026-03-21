package routines

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetRoutineByIDUseCase struct {
	routineRepo       port.RoutineRepository
	routineDetailRepo port.RoutineDetailRepository
}

func NewGetRoutineByIDUseCase(
	routineRepo port.RoutineRepository,
	routineDetailRepo port.RoutineDetailRepository,
) *GetRoutineByIDUseCase {
	return &GetRoutineByIDUseCase{
		routineRepo:       routineRepo,
		routineDetailRepo: routineDetailRepo,
	}
}

func (uc *GetRoutineByIDUseCase) Execute(ctx context.Context, id int) (*dto.RoutineWithDetailsResponse, error) {
	routine, err := uc.routineRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	details, err := uc.routineDetailRepo.GetByRoutineID(ctx, id)
	if err != nil {
		return nil, err
	}

	detailResponses := make([]dto.RoutineDetailResponse, 0, len(details))
	for _, d := range details {
		detailResponses = append(detailResponses, dto.RoutineDetailResponse{
			ID:           d.ID,
			DayOfWeek:    d.DayOfWeek,
			ExerciseID:   d.ExerciseID,
			Sets:         d.Sets,
			Reps:         d.Reps,
			DisplayOrder: d.DisplayOrder,
		})
	}

	return &dto.RoutineWithDetailsResponse{
		ID:         routine.ID,
		CustomerID: routine.CustomerID,
		TemplateID: routine.TemplateID,
		Name:       routine.Name,
		IsActive:   routine.IsActive,
		Details:    detailResponses,
	}, nil
}
