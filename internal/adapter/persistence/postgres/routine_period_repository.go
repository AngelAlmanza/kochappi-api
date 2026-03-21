package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgresRoutinePeriodRepository struct {
	db *gorm.DB
}

func NewPostgresRoutinePeriodRepository(db *gorm.DB) *PostgresRoutinePeriodRepository {
	return &PostgresRoutinePeriodRepository{db: db}
}

func (r *PostgresRoutinePeriodRepository) GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutinePeriod, error) {
	var models []model.RoutinePeriodModel
	if err := r.db.WithContext(ctx).Where("routine_id = ?", routineID).Find(&models).Error; err != nil {
		return nil, err
	}

	periods := make([]entity.RoutinePeriod, 0, len(models))
	for _, m := range models {
		periods = append(periods, *m.ToDomain())
	}
	return periods, nil
}

func (r *PostgresRoutinePeriodRepository) GetOngoingByRoutineID(ctx context.Context, routineID int) (*entity.RoutinePeriod, error) {
	var m model.RoutinePeriodModel
	err := r.db.WithContext(ctx).Where("routine_id = ? AND ended_at IS NULL", routineID).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresRoutinePeriodRepository) Create(ctx context.Context, period *entity.RoutinePeriod) error {
	m := model.RoutinePeriodModelFromDomain(period)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	period.ID = m.ID
	return nil
}

func (r *PostgresRoutinePeriodRepository) Update(ctx context.Context, period *entity.RoutinePeriod) error {
	m := model.RoutinePeriodModelFromDomain(period)
	return r.db.WithContext(ctx).Save(m).Error
}
