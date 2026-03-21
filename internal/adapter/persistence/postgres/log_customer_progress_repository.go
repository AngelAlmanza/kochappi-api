package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type LogCustomerProgressRepository struct {
	db *gorm.DB
}

func NewLogCustomerProgressRepository(db *gorm.DB) *LogCustomerProgressRepository {
	return &LogCustomerProgressRepository{db: db}
}

func (r *LogCustomerProgressRepository) GetByCustomerID(ctx context.Context, customerID int) ([]entity.LogCustomerProgress, error) {
	var models []model.LogCustomerProgressModel
	if err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Order("check_date DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	logs := make([]entity.LogCustomerProgress, 0, len(models))
	for _, m := range models {
		logs = append(logs, *m.ToDomain())
	}
	return logs, nil
}

func (r *LogCustomerProgressRepository) GetByID(ctx context.Context, id int) (*entity.LogCustomerProgress, error) {
	var m model.LogCustomerProgressModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.ProgressLogNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *LogCustomerProgressRepository) Create(ctx context.Context, log *entity.LogCustomerProgress) error {
	m := model.LogCustomerProgressModelFromDomain(log)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	log.ID = m.ID
	return nil
}

func (r *LogCustomerProgressRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.LogCustomerProgressModel{}, id).Error
}
