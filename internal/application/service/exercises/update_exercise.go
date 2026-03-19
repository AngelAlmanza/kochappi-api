package exercises

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
)

type UpdateExerciseUseCase struct {
	exerciseRepo port.ExerciseRepository
}

func NewUpdateExerciseUseCase(exerciseRepo port.ExerciseRepository) *UpdateExerciseUseCase {
	return &UpdateExerciseUseCase{exerciseRepo: exerciseRepo}
}

func (uc *UpdateExerciseUseCase) Execute(ctx context.Context, id int, req *dto.UpdateExerciseRequest) (*dto.ExerciseResponse, error) {
	exercise, err := uc.exerciseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	exercise.Name = req.Name
	exercise.VideoURL = req.VideoURL

	if err := uc.exerciseRepo.Update(ctx, exercise); err != nil {
		return nil, err
	}

	return &dto.ExerciseResponse{
		ID:       exercise.ID,
		Name:     exercise.Name,
		VideoURL: exercise.VideoURL,
	}, nil
}
