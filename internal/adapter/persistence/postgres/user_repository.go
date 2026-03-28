package postgres

import (
	"context"
	"errors"
	"strconv"
	"time"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

func (r *PostgresUserRepository) GetAll(ctx context.Context, role *entity.Role, includeWithCustomers bool) ([]entity.User, error) {
	var models []model.UserModel
	query := r.db.WithContext(ctx)
	if role != nil {
		query = query.Where("role = ?", string(*role))
	}

	if !includeWithCustomers {
		query = query.Where("id NOT IN (SELECT user_id FROM customers)")
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}
	users := make([]entity.User, 0, len(models))
	for _, m := range models {
		users = append(users, *m.ToDomainEntity())
	}
	return users, nil
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	m := model.UserModelFromEntity(user)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	user.ID = m.ID
	return nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	var m model.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.UserNotFoundError{Identifier: strconv.Itoa(id)}
		}
		return nil, err
	}
	return m.ToDomainEntity(), nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var m model.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.UserNotFoundError{Identifier: email}
		}
		return nil, err
	}
	return m.ToDomainEntity(), nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	user.UpdatedAt = time.Now()
	m := model.UserModelFromEntity(user)
	return r.db.WithContext(ctx).Save(m).Error
}
