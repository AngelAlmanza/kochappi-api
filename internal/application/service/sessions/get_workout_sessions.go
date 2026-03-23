package sessions

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetWorkoutSessionsUseCase struct {
	routineRepo        port.RoutineRepository
	workoutSessionRepo port.WorkoutSessionRepository
}

func NewGetWorkoutSessionsUseCase(
	routineRepo port.RoutineRepository,
	workoutSessionRepo port.WorkoutSessionRepository,
) *GetWorkoutSessionsUseCase {
	return &GetWorkoutSessionsUseCase{
		routineRepo:        routineRepo,
		workoutSessionRepo: workoutSessionRepo,
	}
}

func (uc *GetWorkoutSessionsUseCase) Execute(ctx context.Context, routineID int, status *string, from *time.Time, to *time.Time) ([]dto.WorkoutSessionResponse, error) {
	if _, err := uc.routineRepo.GetByID(ctx, routineID); err != nil {
		return nil, err
	}

	sessions, err := uc.workoutSessionRepo.GetByCriteria(ctx, port.WorkoutSessionCriteria{
		RoutineID: &routineID,
		Status:    status,
		DateFrom:  from,
		DateTo:    to,
	})
	if err != nil {
		return nil, err
	}

	responses := make([]dto.WorkoutSessionResponse, 0, len(sessions))
	for _, s := range sessions {
		responses = append(responses, mapWorkoutSessionToResponse(&s))
	}
	return responses, nil
}
