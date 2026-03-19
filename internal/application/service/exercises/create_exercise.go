package exercises

import (
	"context"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type CreateExerciseUseCase struct {
	exerciseRepo port.ExerciseRepository
}

func NewCreateExerciseUseCase(exerciseRepo port.ExerciseRepository) *CreateExerciseUseCase {
	return &CreateExerciseUseCase{exerciseRepo: exerciseRepo}
}

func (uc *CreateExerciseUseCase) Execute(ctx context.Context, req *dto.CreateExerciseRequest) (*dto.ExerciseResponse, error) {
	exercise := entity.NewExercise(req.Name, req.VideoURL)

	if err := uc.exerciseRepo.Create(ctx, exercise); err != nil {
		return nil, err
	}

	return &dto.ExerciseResponse{
		ID:       exercise.ID,
		Name:     exercise.Name,
		VideoURL: exercise.VideoURL,
	}, nil
}
