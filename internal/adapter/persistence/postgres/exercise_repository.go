package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresExerciseRepository struct {
	db *gorm.DB
}

func NewPostgresExerciseRepository(db *gorm.DB) *PostgresExerciseRepository {
	return &PostgresExerciseRepository{db: db}
}

func (r *PostgresExerciseRepository) GetAll(ctx context.Context) ([]entity.Exercise, error) {
	var models []model.ExerciseModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	exercises := make([]entity.Exercise, 0, len(models))
	for _, m := range models {
		exercises = append(exercises, *m.ToDomain())
	}
	return exercises, nil
}

func (r *PostgresExerciseRepository) GetByID(ctx context.Context, id int) (*entity.Exercise, error) {
	var m model.ExerciseModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.ExerciseNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresExerciseRepository) GetByIDs(ctx context.Context, ids []int) ([]entity.Exercise, error) {
	if len(ids) == 0 {
		return []entity.Exercise{}, nil
	}

	var models []model.ExerciseModel
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&models).Error; err != nil {
		return nil, err
	}

	if len(models) != len(ids) {
		found := make(map[int]struct{}, len(models))
		for _, m := range models {
			found[m.ID] = struct{}{}
		}
		for _, id := range ids {
			if _, ok := found[id]; !ok {
				return nil, &domainerror.ExerciseNotFoundError{ID: id}
			}
		}
	}

	exercises := make([]entity.Exercise, 0, len(models))
	for _, m := range models {
		exercises = append(exercises, *m.ToDomain())
	}
	return exercises, nil
}

func (r *PostgresExerciseRepository) Create(ctx context.Context, exercise *entity.Exercise) error {
	m := model.ExerciseModelFromDomain(exercise)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	exercise.ID = m.ID
	return nil
}

func (r *PostgresExerciseRepository) Update(ctx context.Context, exercise *entity.Exercise) error {
	m := model.ExerciseModelFromDomain(exercise)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *PostgresExerciseRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.ExerciseModel{}, id).Error
}
