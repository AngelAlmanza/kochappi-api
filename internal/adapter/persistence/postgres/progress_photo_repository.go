package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type ProgressPhotoRepository struct {
	db *gorm.DB
}

func NewProgressPhotoRepository(db *gorm.DB) *ProgressPhotoRepository {
	return &ProgressPhotoRepository{db: db}
}

func (r *ProgressPhotoRepository) GetByLogID(ctx context.Context, logID int) ([]entity.ProgressPhoto, error) {
	var models []model.ProgressPhotoModel
	if err := r.db.WithContext(ctx).Where("log_customer_progress_id = ?", logID).Find(&models).Error; err != nil {
		return nil, err
	}

	photos := make([]entity.ProgressPhoto, 0, len(models))
	for _, m := range models {
		photos = append(photos, *m.ToDomain())
	}
	return photos, nil
}

func (r *ProgressPhotoRepository) GetByID(ctx context.Context, id int) (*entity.ProgressPhoto, error) {
	var m model.ProgressPhotoModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.ProgressPhotoNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *ProgressPhotoRepository) Create(ctx context.Context, photo *entity.ProgressPhoto) error {
	m := model.ProgressPhotoModelFromDomain(photo)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	photo.ID = m.ID
	return nil
}

func (r *ProgressPhotoRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.ProgressPhotoModel{}, id).Error
}

func (r *ProgressPhotoRepository) DeleteByLogID(ctx context.Context, logID int) error {
	return r.db.WithContext(ctx).Where("log_customer_progress_id = ?", logID).Delete(&model.ProgressPhotoModel{}).Error
}
