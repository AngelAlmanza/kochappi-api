package model

import (
	"time"

	"kochappi/internal/domain/entity"
)

type ExerciseRoutineModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	ExerciseID int       `gorm:"type:integer;not null;index"`
	RoutineID  int       `gorm:"type:integer;not null;index"`
	CreatedAt  time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}

func (ExerciseRoutineModel) TableName() string {
	return "exercise_routine"
}

func (m *ExerciseRoutineModel) ToDomain() *entity.ExerciseRoutine {
	return &entity.ExerciseRoutine{
		ID:         m.ID,
		ExerciseID: m.ExerciseID,
		RoutineID:  m.RoutineID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func ExerciseRoutineModelFromDomain(er *entity.ExerciseRoutine) *ExerciseRoutineModel {
	return &ExerciseRoutineModel{
		ID:         er.ID,
		ExerciseID: er.ExerciseID,
		RoutineID:  er.RoutineID,
		CreatedAt:  er.CreatedAt,
		UpdatedAt:  er.UpdatedAt,
	}
}
