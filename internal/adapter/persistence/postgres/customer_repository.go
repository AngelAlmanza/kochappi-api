package postgres

import (
	"context"
	"errors"

	"kochappi/internal/adapter/persistence/postgres/model"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll(ctx context.Context) ([]entity.Customer, error) {
	var models []model.CustomerModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	customers := make([]entity.Customer, 0, len(models))
	for _, m := range models {
		customers = append(customers, *m.ToDomain())
	}
	return customers, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*entity.Customer, error) {
	var m model.CustomerModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.CustomerNotFoundError{ID: id}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *CustomerRepository) GetByUserID(ctx context.Context, userID int) (*entity.Customer, error) {
	var m model.CustomerModel
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domainerror.UserNotCustomerError{ID: userID}
		}
		return nil, err
	}
	return m.ToDomain(), nil
}

func (r *CustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	m := model.CustomerModelFromDomain(customer)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	customer.ID = m.ID
	return nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	m := model.CustomerModelFromDomain(customer)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.CustomerModel{}, id).Error
}
