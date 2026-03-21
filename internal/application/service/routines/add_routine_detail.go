package routines

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type AddRoutineDetailUseCase struct {
	routineRepo       port.RoutineRepository
	routineDetailRepo port.RoutineDetailRepository
	exerciseRepo      port.ExerciseRepository
}

func NewAddRoutineDetailUseCase(
	routineRepo port.RoutineRepository,
	routineDetailRepo port.RoutineDetailRepository,
	exerciseRepo port.ExerciseRepository,
) *AddRoutineDetailUseCase {
	return &AddRoutineDetailUseCase{
		routineRepo:       routineRepo,
		routineDetailRepo: routineDetailRepo,
		exerciseRepo:      exerciseRepo,
	}
}

func (uc *AddRoutineDetailUseCase) Execute(ctx context.Context, routineID int, req *dto.AddRoutineDetailRequest) (*dto.RoutineDetailResponse, error) {
	if _, err := uc.routineRepo.GetByID(ctx, routineID); err != nil {
		return nil, err
	}

	if _, err := uc.exerciseRepo.GetByID(ctx, req.ExerciseID); err != nil {
		return nil, err
	}

	detail := entity.NewRoutineDetail(routineID, req.DayOfWeek, req.ExerciseID, req.Sets, req.Reps, req.DisplayOrder)
	if err := uc.routineDetailRepo.Create(ctx, detail); err != nil {
		return nil, err
	}

	return &dto.RoutineDetailResponse{
		ID:           detail.ID,
		DayOfWeek:    detail.DayOfWeek,
		ExerciseID:   detail.ExerciseID,
		Sets:         detail.Sets,
		Reps:         detail.Reps,
		DisplayOrder: detail.DisplayOrder,
	}, nil
}
