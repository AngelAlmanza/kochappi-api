package sessions

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type UpdateExerciseLogUseCase struct {
	logExerciseSessionRepo port.LogExerciseSessionRepository
}

func NewUpdateExerciseLogUseCase(logExerciseSessionRepo port.LogExerciseSessionRepository) *UpdateExerciseLogUseCase {
	return &UpdateExerciseLogUseCase{logExerciseSessionRepo: logExerciseSessionRepo}
}

func (uc *UpdateExerciseLogUseCase) Execute(ctx context.Context, id int, req *dto.UpdateExerciseLogRequest) (*dto.ExerciseLogResponse, error) {
	log, err := uc.logExerciseSessionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	log.SetNumber = req.SetNumber
	log.RepsDone = req.RepsDone
	log.Weight = req.Weight
	log.Notes = req.Notes
	log.UpdatedAt = time.Now()

	if err := uc.logExerciseSessionRepo.Update(ctx, log); err != nil {
		return nil, err
	}

	resp := mapExerciseLogToResponse(log)
	return &resp, nil
}
