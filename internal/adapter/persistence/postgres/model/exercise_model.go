package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type ExerciseModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	VideoURL  string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (ExerciseModel) TableName() string {
	return "exercises"
}

func (m *ExerciseModel) ToDomain() *entity.Exercise {
	return &entity.Exercise{
		ID:        m.ID,
		Name:      m.Name,
		VideoURL:  m.VideoURL,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ExerciseModelFromDomain(e *entity.Exercise) *ExerciseModel {
	return &ExerciseModel{
		ID:        e.ID,
		Name:      e.Name,
		VideoURL:  e.VideoURL,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
