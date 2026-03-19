package postgres

import (
	"kochappi/internal/adapter/persistence/postgres/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.UserModel{},
		&model.RefreshTokenModel{},
	)
}
