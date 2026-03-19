package postgres

import (
	"context"
	"time"

	"kochappi/internal/adapter/persistence/postgres/model"

	"gorm.io/gorm"
)

type PostgresRefreshTokenRepository struct {
	db *gorm.DB
}

func NewPostgresRefreshTokenRepository(db *gorm.DB) *PostgresRefreshTokenRepository {
	return &PostgresRefreshTokenRepository{db: db}
}

func (r *PostgresRefreshTokenRepository) Store(ctx context.Context, userID string, tokenID string, expiresAt int64) error {
	m := &model.RefreshTokenModel{
		ID:        tokenID,
		UserID:    userID,
		ExpiresAt: time.Unix(expiresAt, 0),
		CreatedAt: time.Now(),
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *PostgresRefreshTokenRepository) Exists(ctx context.Context, tokenID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.RefreshTokenModel{}).
		Where("id = ? AND expires_at > ?", tokenID, time.Now()).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PostgresRefreshTokenRepository) DeleteByID(ctx context.Context, tokenID string) error {
	return r.db.WithContext(ctx).
		Where("id = ?", tokenID).
		Delete(&model.RefreshTokenModel{}).Error
}

func (r *PostgresRefreshTokenRepository) DeleteAllByUserID(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&model.RefreshTokenModel{}).Error
}
