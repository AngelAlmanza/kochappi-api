package sessions

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type GetWorkoutSessionByIDUseCase struct {
	workoutSessionRepo     port.WorkoutSessionRepository
	logExerciseSessionRepo port.LogExerciseSessionRepository
}

func NewGetWorkoutSessionByIDUseCase(
	workoutSessionRepo port.WorkoutSessionRepository,
	logExerciseSessionRepo port.LogExerciseSessionRepository,
) *GetWorkoutSessionByIDUseCase {
	return &GetWorkoutSessionByIDUseCase{
		workoutSessionRepo:     workoutSessionRepo,
		logExerciseSessionRepo: logExerciseSessionRepo,
	}
}

func (uc *GetWorkoutSessionByIDUseCase) Execute(ctx context.Context, id int) (*dto.WorkoutSessionWithLogsResponse, error) {
	session, err := uc.workoutSessionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	logs, err := uc.logExerciseSessionRepo.GetByWorkoutSessionID(ctx, session.ID)
	if err != nil {
		return nil, err
	}

	logResponses := make([]dto.ExerciseLogResponse, 0, len(logs))
	for _, l := range logs {
		logResponses = append(logResponses, mapExerciseLogToResponse(&l))
	}

	resp := mapWorkoutSessionToResponse(session)
	return &dto.WorkoutSessionWithLogsResponse{
		ID:           resp.ID,
		RoutineID:    resp.RoutineID,
		ScheduledDay: resp.ScheduledDay,
		ActualDate:   resp.ActualDate,
		Status:       resp.Status,
		StartedAt:    resp.StartedAt,
		FinishedAt:   resp.FinishedAt,
		ExerciseLogs: logResponses,
	}, nil
}
