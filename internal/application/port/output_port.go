package port

import (
	"context"
	"io"

	"kochappi/internal/domain/entity"
)

// Users and auth

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}

type RefreshTokenRepository interface {
	Store(ctx context.Context, userID int, tokenID string, expiresAt int64) error
	Exists(ctx context.Context, tokenID string) (bool, error)
	DeleteByID(ctx context.Context, tokenID string) error
	DeleteAllByUserID(ctx context.Context, userID int) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type TokenProvider interface {
	GenerateAccessToken(userID int, role string) (string, error)
	GenerateRefreshToken(userID int) (tokenString string, tokenID string, expiresAt int64, err error)
	ValidateAccessToken(tokenString string) (userID int, role string, err error)
	ValidateRefreshToken(tokenString string) (userID int, tokenID string, err error)
}

type OTPService interface {
	GenerateCode() string
	Send(ctx context.Context, email string, code string) error
}

// Customers

type CustomerRepository interface {
	GetAll(ctx context.Context) ([]entity.Customer, error)
	GetByID(ctx context.Context, id int) (*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) error
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id int) error
	GetByUserID(ctx context.Context, userID int) (*entity.Customer, error)
}

// Templates

type TemplateRepository interface {
	GetAll(ctx context.Context) ([]entity.Template, error)
	GetByID(ctx context.Context, id int) (*entity.Template, error)
	Create(ctx context.Context, template *entity.Template) error
	Update(ctx context.Context, template *entity.Template) error
	Delete(ctx context.Context, id int) error
}

type TemplateDetailRepository interface {
	GetByTemplateID(ctx context.Context, templateID int) ([]entity.TemplateDetail, error)
	GetByID(ctx context.Context, id int) (*entity.TemplateDetail, error)
	Create(ctx context.Context, detail *entity.TemplateDetail) error
	CreateBulk(ctx context.Context, details []*entity.TemplateDetail) error
	DeleteByID(ctx context.Context, id int) error
}

// Routines

type RoutineRepository interface {
	GetAll(ctx context.Context) ([]entity.Routine, error)
	GetByID(ctx context.Context, id int) (*entity.Routine, error)
	GetByCustomerID(ctx context.Context, customerID int) ([]entity.Routine, error)
	GetActiveByCustomerID(ctx context.Context, customerID int) (*entity.Routine, error)
	Create(ctx context.Context, routine *entity.Routine) error
	Update(ctx context.Context, routine *entity.Routine) error
	GetAllActive(ctx context.Context) ([]entity.Routine, error)
}

type RoutineDetailRepository interface {
	GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutineDetail, error)
	GetByID(ctx context.Context, id int) (*entity.RoutineDetail, error)
	Create(ctx context.Context, detail *entity.RoutineDetail) error
	CreateBulk(ctx context.Context, details []*entity.RoutineDetail) error
	DeleteByID(ctx context.Context, id int) error
}

type RoutinePeriodRepository interface {
	GetByRoutineID(ctx context.Context, routineID int) ([]entity.RoutinePeriod, error)
	GetOngoingByRoutineID(ctx context.Context, routineID int) (*entity.RoutinePeriod, error)
	Create(ctx context.Context, period *entity.RoutinePeriod) error
	Update(ctx context.Context, period *entity.RoutinePeriod) error
}

// Exercises

type ExerciseRepository interface {
	GetAll(ctx context.Context) ([]entity.Exercise, error)
	GetByID(ctx context.Context, id int) (*entity.Exercise, error)
	GetByIDs(ctx context.Context, ids []int) ([]entity.Exercise, error)
	Create(ctx context.Context, exercise *entity.Exercise) error
	Update(ctx context.Context, exercise *entity.Exercise) error
	Delete(ctx context.Context, id int) error
}

// Progress

type LogCustomerProgressRepository interface {
	GetByCustomerID(ctx context.Context, customerID int) ([]entity.LogCustomerProgress, error)
	GetByID(ctx context.Context, id int) (*entity.LogCustomerProgress, error)
	Create(ctx context.Context, log *entity.LogCustomerProgress) error
	Delete(ctx context.Context, id int) error
}

type ProgressPhotoRepository interface {
	GetByLogID(ctx context.Context, logID int) ([]entity.ProgressPhoto, error)
	GetByID(ctx context.Context, id int) (*entity.ProgressPhoto, error)
	Create(ctx context.Context, photo *entity.ProgressPhoto) error
	Delete(ctx context.Context, id int) error
	DeleteByLogID(ctx context.Context, logID int) error
}

// Storage

type FileStorage interface {
	Upload(ctx context.Context, filename string, file io.Reader) (url string, err error)
	Delete(ctx context.Context, url string) error
}
