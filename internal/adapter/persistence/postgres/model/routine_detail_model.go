package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type RoutineDetailModel struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	RoutineID    int       `gorm:"type:integer;not null;index"`
	DayOfWeek    int16     `gorm:"type:smallint;not null"`
	ExerciseID   int       `gorm:"type:integer;not null"`
	Sets         int16     `gorm:"type:smallint;not null"`
	Reps         int16     `gorm:"type:smallint;not null"`
	DisplayOrder int16     `gorm:"type:smallint;not null"`
	CreatedAt    time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (RoutineDetailModel) TableName() string {
	return "routine_detail"
}

func (m *RoutineDetailModel) ToDomain() *entity.RoutineDetail {
	return &entity.RoutineDetail{
		ID:           m.ID,
		RoutineID:    m.RoutineID,
		DayOfWeek:    m.DayOfWeek,
		ExerciseID:   m.ExerciseID,
		Sets:         m.Sets,
		Reps:         m.Reps,
		DisplayOrder: m.DisplayOrder,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func RoutineDetailModelFromDomain(rd *entity.RoutineDetail) *RoutineDetailModel {
	return &RoutineDetailModel{
		ID:           rd.ID,
		RoutineID:    rd.RoutineID,
		DayOfWeek:    rd.DayOfWeek,
		ExerciseID:   rd.ExerciseID,
		Sets:         rd.Sets,
		Reps:         rd.Reps,
		DisplayOrder: rd.DisplayOrder,
		CreatedAt:    rd.CreatedAt,
		UpdatedAt:    rd.UpdatedAt,
	}
}
