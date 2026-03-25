package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	"kochappi/internal/shared/logger"

	"gorm.io/gorm"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SeedConfig struct {
	TrainerName     string
	TrainerEmail    string
	TrainerPassword string
}

func Seed(db *gorm.DB, hasher PasswordHasher, cfg SeedConfig) error {
	ctx := context.Background()
	return seedTrainer(ctx, db, hasher, cfg)
}

func seedTrainer(ctx context.Context, db *gorm.DB, hasher PasswordHasher, cfg SeedConfig) error {
	var existing model.UserModel
	err := db.WithContext(ctx).Where("email = ?", cfg.TrainerEmail).First(&existing).Error
	if err == nil {
		logger.Info.Printf("Seeder: trainer user '%s' already exists, skipping", cfg.TrainerEmail)
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := hasher.Hash(cfg.TrainerPassword)
	if err != nil {
		return err
	}

	user := entity.NewUser(cfg.TrainerName, cfg.TrainerEmail, hash, entity.ROLE_TRAINER)
	m := model.UserModelFromEntity(user)
	if err := db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}

	logger.Info.Printf("Seeder: trainer user '%s' created successfully", cfg.TrainerEmail)
	return nil
}
