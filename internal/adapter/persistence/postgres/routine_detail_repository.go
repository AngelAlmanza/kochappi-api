package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type PostgresRoutineDetailRepository struct {
	db *gorm.DB
}

func NewPostgresRoutineDetailRepository(db *gorm.DB) *PostgresRoutineDetailRepository {
	return &PostgresRoutineDetailRepository{db: db}
}

func (r *PostgresRoutineDetailRepository) GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutineDetail, error) {
	var models []model.RoutineDetailModel
	if err := r.db.WithContext(ctx).Where("routine_id = ?", routineID).Order("display_order ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	details := make([]entity.RoutineDetail, 0, len(models))
	for _, m := range models {
		details = append(details, *m.ToDomain())
	}
	return details, nil
}

func (r *PostgresRoutineDetailRepository) GetByID(ctx context.Context, id int) (*entity.RoutineDetail, error) {
	var m model.RoutineDetailModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.RoutineDetailNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *PostgresRoutineDetailRepository) Create(ctx context.Context, detail *entity.RoutineDetail) error {
	m := model.RoutineDetailModelFromDomain(detail)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	detail.ID = m.ID
	return nil
}

func (r *PostgresRoutineDetailRepository) CreateBulk(ctx context.Context, details []*entity.RoutineDetail) error {
	if len(details) == 0 {
		return nil
	}

	models := make([]*model.RoutineDetailModel, 0, len(details))
	for _, d := range details {
		models = append(models, model.RoutineDetailModelFromDomain(d))
	}

	if err := r.db.WithContext(ctx).Create(&models).Error; err != nil {
		return err
	}

	for i, m := range models {
		details[i].ID = m.ID
	}
	return nil
}

func (r *PostgresRoutineDetailRepository) DeleteByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.RoutineDetailModel{}, id).Error
}
