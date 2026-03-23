package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresLogExerciseSessionRepository struct {
	db *gorm.DB
}

func NewPostgresLogExerciseSessionRepository(db *gorm.DB) *PostgresLogExerciseSessionRepository {
	return &PostgresLogExerciseSessionRepository{db: db}
}

func (r *PostgresLogExerciseSessionRepository) GetByID(ctx context.Context, id int) (*entity.LogExerciseSession, error) {
	var m model.LogExerciseSessionModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.LogExerciseSessionNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresLogExerciseSessionRepository) GetByWorkoutSessionID(ctx context.Context, workoutSessionID int) ([]entity.LogExerciseSession, error) {
	var models []model.LogExerciseSessionModel
	if err := r.db.WithContext(ctx).Where("workout_session_id = ?", workoutSessionID).Order("set_number ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	logs := make([]entity.LogExerciseSession, 0, len(models))
	for _, m := range models {
		logs = append(logs, *m.ToDomain())
	}
	return logs, nil
}

func (r *PostgresLogExerciseSessionRepository) Create(ctx context.Context, log *entity.LogExerciseSession) error {
	m := model.LogExerciseSessionModelFromDomain(log)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	log.ID = m.ID
	return nil
}

func (r *PostgresLogExerciseSessionRepository) Update(ctx context.Context, log *entity.LogExerciseSession) error {
	m := model.LogExerciseSessionModelFromDomain(log)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *PostgresLogExerciseSessionRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.LogExerciseSessionModel{}, id).Error
}
