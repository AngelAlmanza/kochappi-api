package postgres

import (
	"kochappi/internal/adapter/persistence/postgres/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.CustomerModel{},
		&model.ExerciseModel{},
		&model.ExerciseRoutineModel{},
		&model.LogCustomerProgressModel{},
		&model.LogExerciseSessionModel{},
		&model.ProgressPhotoModel{},
		&model.RefreshTokenModel{},
		&model.RoutineDetailModel{},
		&model.RoutinePeriodModel{},
		&model.RoutineModel{},
		&model.TemplateDetailModel{},
		&model.TemplateModel{},
		&model.UserModel{},
		&model.WorkoutSessionModel{},
	)
}
