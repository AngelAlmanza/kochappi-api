package sessions

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type UpdateWorkoutSessionStatusUseCase struct {
	workoutSessionRepo port.WorkoutSessionRepository
}

func NewUpdateWorkoutSessionStatusUseCase(workoutSessionRepo port.WorkoutSessionRepository) *UpdateWorkoutSessionStatusUseCase {
	return &UpdateWorkoutSessionStatusUseCase{workoutSessionRepo: workoutSessionRepo}
}

func (uc *UpdateWorkoutSessionStatusUseCase) Execute(ctx context.Context, id int, status string) (*dto.WorkoutSessionResponse, error) {
	session, err := uc.workoutSessionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := session.TransitionTo(entity.WorkoutStatus(status)); err != nil {
		return nil, err
	}

	if err := uc.workoutSessionRepo.Update(ctx, session); err != nil {
		return nil, err
	}

	resp := mapWorkoutSessionToResponse(session)
	return &resp, nil
}
