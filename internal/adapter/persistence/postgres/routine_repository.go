package postgres

import (
	"context"
	"errors"
	"time"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/shared/logger"

	"gorm.io/gorm"
)

type PostgresRoutineRepository struct {
	db *gorm.DB
}

func NewPostgresRoutineRepository(db *gorm.DB) *PostgresRoutineRepository {
	return &PostgresRoutineRepository{db: db}
}

func (r *PostgresRoutineRepository) GetAll(ctx context.Context) ([]entity.Routine, error) {
	var models []model.RoutineModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	routines := make([]entity.Routine, 0, len(models))
	for _, m := range models {
		routines = append(routines, *m.ToDomain())
	}
	return routines, nil
}

func (r *PostgresRoutineRepository) GetByID(ctx context.Context, id int) (*entity.Routine, error) {
	var m model.RoutineModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.RoutineNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresRoutineRepository) GetByCustomerID(ctx context.Context, customerID int) ([]entity.Routine, error) {
	var models []model.RoutineModel
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Find(&models).Error; err != nil {
		return nil, err
	}

	routines := make([]entity.Routine, 0, len(models))
	for _, m := range models {
		routines = append(routines, *m.ToDomain())
	}
	return routines, nil
}

func (r *PostgresRoutineRepository) GetActiveByCustomerID(ctx context.Context, customerID int) (*entity.Routine, error) {
	var m model.RoutineModel
	err := r.db.WithContext(ctx).Where("customer_id = ? AND is_active = true", customerID).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresRoutineRepository) Create(ctx context.Context, routine *entity.Routine) error {
	m := model.RoutineModelFromDomain(routine)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	routine.ID = m.ID
	return nil
}

func (r *PostgresRoutineRepository) Update(ctx context.Context, routine *entity.Routine) error {
	m := model.RoutineModelFromDomain(routine)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *PostgresRoutineRepository) GetAllActive(ctx context.Context) ([]entity.Routine, error) {
	var models []model.RoutineModel
	if err := r.db.WithContext(ctx).Where("is_active = true").Find(&models).Error; err != nil {
		return nil, err
	}

	routines := make([]entity.Routine, 0, len(models))
	for _, m := range models {
		routines = append(routines, *m.ToDomain())
	}
	return routines, nil
}

func (r *PostgresRoutineRepository) GetRoutinesToGenerateSessions(ctx context.Context, dayOfWeek int16, date time.Time) ([]entity.Routine, error) {
	logger.Info.Printf("Get Routines to Generate Sessions for day: %d and date: %s", dayOfWeek, date)
	var models []model.RoutineModel
	err := r.db.WithContext(ctx).Raw(`
		SELECT r.*
		FROM routines r
		JOIN routine_details rd ON r.id = rd.routine_id
		WHERE r.is_active = true
		AND rd.day_of_week = ?
		AND NOT EXISTS (
			SELECT 1
			FROM workout_sessions ws
			WHERE ws.routine_id = r.id
			AND ws.actual_date = ?
		)
		GROUP BY r.id
	`, dayOfWeek, date).Scan(&models).Error
	if err != nil {
		return nil, err
	}

	routines := make([]entity.Routine, 0, len(models))
	for _, m := range models {
		routines = append(routines, *m.ToDomain())
	}
	return routines, nil
}
