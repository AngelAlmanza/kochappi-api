package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresWorkoutSessionRepository struct {
	db *gorm.DB
}

func NewPostgresWorkoutSessionRepository(db *gorm.DB) *PostgresWorkoutSessionRepository {
	return &PostgresWorkoutSessionRepository{db: db}
}

func (r *PostgresWorkoutSessionRepository) GetByID(ctx context.Context, id int) (*entity.WorkoutSession, error) {
	var m model.WorkoutSessionModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.WorkoutSessionNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresWorkoutSessionRepository) GetByCriteria(ctx context.Context, criteria port.WorkoutSessionCriteria) ([]entity.WorkoutSession, error) {
	var models []model.WorkoutSessionModel

	q := r.db.WithContext(ctx).Model(&model.WorkoutSessionModel{})

	if criteria.RoutineID != nil {
		q = q.Where("routine_id = ?", *criteria.RoutineID)
	}
	if criteria.Status != nil {
		q = q.Where("status = ?", *criteria.Status)
	}
	if criteria.Date != nil {
		q = q.Where("actual_date = ?", *criteria.Date)
	} else {
		if criteria.DateFrom != nil {
			q = q.Where("actual_date >= ?", *criteria.DateFrom)
		}
		if criteria.DateTo != nil {
			q = q.Where("actual_date <= ?", *criteria.DateTo)
		}
	}

	if err := q.Order("actual_date DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	sessions := make([]entity.WorkoutSession, 0, len(models))
	for _, m := range models {
		sessions = append(sessions, *m.ToDomain())
	}
	return sessions, nil
}

func (r *PostgresWorkoutSessionRepository) Create(ctx context.Context, session *entity.WorkoutSession) error {
	m := model.WorkoutSessionModelFromDomain(session)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	session.ID = m.ID
	return nil
}

func (r *PostgresWorkoutSessionRepository) CreateBulk(ctx context.Context, sessions []*entity.WorkoutSession) error {
	if len(sessions) == 0 {
		return nil
	}

	models := make([]*model.WorkoutSessionModel, 0, len(sessions))
	for _, s := range sessions {
		models = append(models, model.WorkoutSessionModelFromDomain(s))
	}

	if err := r.db.WithContext(ctx).Create(&models).Error; err != nil {
		return err
	}

	for i, m := range models {
		sessions[i].ID = m.ID
	}
	return nil
}

func (r *PostgresWorkoutSessionRepository) Update(ctx context.Context, session *entity.WorkoutSession) error {
	m := model.WorkoutSessionModelFromDomain(session)
	return r.db.WithContext(ctx).Save(m).Error
}
