package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type LogExerciseSessionModel struct {
	ID               int       `gorm:"primaryKey;autoIncrement"`
	WorkoutSessionID int       `gorm:"type:integer;not null;index"`
	RoutineDetailID  int       `gorm:"type:integer;not null"`
	SetNumber        int16     `gorm:"type:smallint;not null"`
	RepsDone         int16     `gorm:"type:smallint;not null"`
	Weight           float64   `gorm:"type:decimal;not null"`
	Notes            string    `gorm:"type:varchar(255)"`
	CreatedAt        time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (LogExerciseSessionModel) TableName() string {
	return "log_exercise_sessions"
}

func (m *LogExerciseSessionModel) ToDomain() *entity.LogExerciseSession {
	return &entity.LogExerciseSession{
		ID:               m.ID,
		WorkoutSessionID: m.WorkoutSessionID,
		RoutineDetailID:  m.RoutineDetailID,
		SetNumber:        m.SetNumber,
		RepsDone:         m.RepsDone,
		Weight:           m.Weight,
		Notes:            m.Notes,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}

func LogExerciseSessionModelFromDomain(l *entity.LogExerciseSession) *LogExerciseSessionModel {
	return &LogExerciseSessionModel{
		ID:               l.ID,
		WorkoutSessionID: l.WorkoutSessionID,
		RoutineDetailID:  l.RoutineDetailID,
		SetNumber:        l.SetNumber,
		RepsDone:         l.RepsDone,
		Weight:           l.Weight,
		Notes:            l.Notes,
		CreatedAt:        l.CreatedAt,
		UpdatedAt:        l.UpdatedAt,
	}
}
