package sessions

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

type CreateExerciseLogUseCase struct {
	workoutSessionRepo     port.WorkoutSessionRepository
	logExerciseSessionRepo port.LogExerciseSessionRepository
	routineDetailRepo      port.RoutineDetailRepository
}

func NewCreateExerciseLogUseCase(
	workoutSessionRepo port.WorkoutSessionRepository,
	logExerciseSessionRepo port.LogExerciseSessionRepository,
	routineDetailRepo port.RoutineDetailRepository,
) *CreateExerciseLogUseCase {
	return &CreateExerciseLogUseCase{
		workoutSessionRepo:     workoutSessionRepo,
		logExerciseSessionRepo: logExerciseSessionRepo,
		routineDetailRepo:      routineDetailRepo,
	}
}

func (uc *CreateExerciseLogUseCase) Execute(ctx context.Context, workoutSessionID int, req *dto.CreateExerciseLogRequest) (*dto.ExerciseLogResponse, error) {
	session, err := uc.workoutSessionRepo.GetByID(ctx, workoutSessionID)
	if err != nil {
		return nil, err
	}

	if session.Status != entity.WorkoutStatusInProgress {
		return nil, &domainerror.InvalidSessionStatusTransitionError{
			From: string(session.Status),
			To:   "log_exercise",
		}
	}

	if _, err := uc.routineDetailRepo.GetByID(ctx, req.RoutineDetailID); err != nil {
		return nil, err
	}

	log := entity.NewLogExerciseSession(
		workoutSessionID,
		req.RoutineDetailID,
		req.SetNumber,
		req.RepsDone,
		req.Weight,
		req.Notes,
	)

	if err := uc.logExerciseSessionRepo.Create(ctx, log); err != nil {
		return nil, err
	}

	resp := mapExerciseLogToResponse(log)
	return &resp, nil
}
