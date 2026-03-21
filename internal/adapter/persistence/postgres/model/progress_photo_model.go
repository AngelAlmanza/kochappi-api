package model

import (
	"time"

	"kochappi/internal/domain/entity"
	"kochappi/internal/domain/value_object"
)

type ProgressPhotoModel struct {
	ID                    int       `gorm:"primaryKey;autoIncrement"`
	URL                   string    `gorm:"type:varchar(255);not null"`
	PictureType           string    `gorm:"type:varchar(255);not null"`
	LogCustomerProgressID int       `gorm:"type:integer;not null;index"`
	CreatedAt             time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt             time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (ProgressPhotoModel) TableName() string {
	return "progress_photo"
}

func (m *ProgressPhotoModel) ToDomain() *entity.ProgressPhoto {
	return &entity.ProgressPhoto{
		ID:                    m.ID,
		URL:                   m.URL,
		PictureType:           value_object.PictureType(m.PictureType),
		LogCustomerProgressID: m.LogCustomerProgressID,
		CreatedAt:             m.CreatedAt,
		UpdatedAt:             m.UpdatedAt,
	}
}

func ProgressPhotoModelFromDomain(p *entity.ProgressPhoto) *ProgressPhotoModel {
	return &ProgressPhotoModel{
		ID:                    p.ID,
		URL:                   p.URL,
		PictureType:           p.PictureType.String(),
		LogCustomerProgressID: p.LogCustomerProgressID,
		CreatedAt:             p.CreatedAt,
		UpdatedAt:             p.UpdatedAt,
	}
}
