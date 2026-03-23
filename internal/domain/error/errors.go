package error

import "fmt"

type DomainError interface {
	error
	Code() string
	IsUserError() bool
}

type UserNotFoundError struct {
	Identifier string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user %s not found", e.Identifier)
}

func (e *UserNotFoundError) Code() string {
	return "USER_NOT_FOUND"
}

func (e *UserNotFoundError) IsUserError() bool {
	return true
}

type EmailAlreadyExistsError struct {
	Email string
}

func (e *EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("email %s is already registered", e.Email)
}

func (e *EmailAlreadyExistsError) Code() string {
	return "EMAIL_ALREADY_EXISTS"
}

func (e *EmailAlreadyExistsError) IsUserError() bool {
	return true
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid email or password"
}

func (e *InvalidCredentialsError) Code() string {
	return "INVALID_CREDENTIALS"
}

func (e *InvalidCredentialsError) IsUserError() bool {
	return true
}

type InvalidOTPError struct{}

func (e *InvalidOTPError) Error() string {
	return "invalid or expired OTP code"
}

func (e *InvalidOTPError) Code() string {
	return "INVALID_OTP"
}

func (e *InvalidOTPError) IsUserError() bool {
	return true
}

type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return "invalid or expired token"
}

func (e *InvalidTokenError) Code() string {
	return "INVALID_TOKEN"
}

func (e *InvalidTokenError) IsUserError() bool {
	return true
}

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "unauthorized access"
}

func (e *UnauthorizedError) Code() string {
	return "UNAUTHORIZED"
}

func (e *UnauthorizedError) IsUserError() bool {
	return true
}

type UserNotCustomerError struct {
	ID int
}

func (e *UserNotCustomerError) Error() string {
	return fmt.Sprintf("user %d is not a customer", e.ID)
}

func (e *UserNotCustomerError) Code() string {
	return "USER_NOT_CUSTOMER"
}

func (e *UserNotCustomerError) IsUserError() bool {
	return true
}

type ExerciseNotFoundError struct {
	ID int
}

func (e *ExerciseNotFoundError) Error() string {
	return fmt.Sprintf("exercise %d not found", e.ID)
}

func (e *ExerciseNotFoundError) Code() string {
	return "EXERCISE_NOT_FOUND"
}

func (e *ExerciseNotFoundError) IsUserError() bool {
	return true
}

type InvalidBirthdateError struct {
	Birthdate string
}

func (e *InvalidBirthdateError) Error() string {
	return fmt.Sprintf("invalid birthdate: %s", e.Birthdate)
}

func (e *InvalidBirthdateError) Code() string {
	return "INVALID_BIRTHDATE"
}

func (e *InvalidBirthdateError) IsUserError() bool {
	return true
}

type CustomerNotFoundError struct {
	ID int
}

func (e *CustomerNotFoundError) Error() string {
	return fmt.Sprintf("customer %d not found", e.ID)
}

func (e *CustomerNotFoundError) Code() string {
	return "CUSTOMER_NOT_FOUND"
}

func (e *CustomerNotFoundError) IsUserError() bool {
	return true
}

type TemplateNotFoundError struct {
	ID int
}

func (e *TemplateNotFoundError) Error() string {
	return fmt.Sprintf("template %d not found", e.ID)
}

func (e *TemplateNotFoundError) Code() string {
	return "TEMPLATE_NOT_FOUND"
}

func (e *TemplateNotFoundError) IsUserError() bool {
	return true
}

type TemplateDetailNotFoundError struct {
	ID int
}

func (e *TemplateDetailNotFoundError) Error() string {
	return fmt.Sprintf("template detail %d not found", e.ID)
}

func (e *TemplateDetailNotFoundError) Code() string {
	return "TEMPLATE_DETAIL_NOT_FOUND"
}

func (e *TemplateDetailNotFoundError) IsUserError() bool {
	return true
}

type CustomerAlreadyExistsError struct {
	UserID int
}

func (e *CustomerAlreadyExistsError) Error() string {
	return fmt.Sprintf("a customer already exists for user %d", e.UserID)
}

func (e *CustomerAlreadyExistsError) Code() string {
	return "CUSTOMER_ALREADY_EXISTS"
}

func (e *CustomerAlreadyExistsError) IsUserError() bool {
	return true
}

type RoutineNotFoundError struct {
	ID int
}

func (e *RoutineNotFoundError) Error() string {
	return fmt.Sprintf("routine %d not found", e.ID)
}

func (e *RoutineNotFoundError) Code() string {
	return "ROUTINE_NOT_FOUND"
}

func (e *RoutineNotFoundError) IsUserError() bool {
	return true
}

type RoutineDetailNotFoundError struct {
	ID int
}

func (e *RoutineDetailNotFoundError) Error() string {
	return fmt.Sprintf("routine detail %d not found", e.ID)
}

func (e *RoutineDetailNotFoundError) Code() string {
	return "ROUTINE_DETAIL_NOT_FOUND"
}

func (e *RoutineDetailNotFoundError) IsUserError() bool {
	return true
}

type ActiveRoutineExistsError struct {
	CustomerID int
}

func (e *ActiveRoutineExistsError) Error() string {
	return fmt.Sprintf("customer %d already has an active routine", e.CustomerID)
}

func (e *ActiveRoutineExistsError) Code() string {
	return "ACTIVE_ROUTINE_EXISTS"
}

func (e *ActiveRoutineExistsError) IsUserError() bool {
	return true
}

type ProgressLogNotFoundError struct {
	ID int
}

func (e *ProgressLogNotFoundError) Error() string {
	return fmt.Sprintf("progress log %d not found", e.ID)
}

func (e *ProgressLogNotFoundError) Code() string {
	return "PROGRESS_LOG_NOT_FOUND"
}

func (e *ProgressLogNotFoundError) IsUserError() bool {
	return true
}

type ProgressPhotoNotFoundError struct {
	ID int
}

func (e *ProgressPhotoNotFoundError) Error() string {
	return fmt.Sprintf("progress photo %d not found", e.ID)
}

func (e *ProgressPhotoNotFoundError) Code() string {
	return "PROGRESS_PHOTO_NOT_FOUND"
}

func (e *ProgressPhotoNotFoundError) IsUserError() bool {
	return true
}

type InvalidCheckDateError struct {
	CheckDate string
}

func (e *InvalidCheckDateError) Error() string {
	return fmt.Sprintf("invalid check date: %s", e.CheckDate)
}

func (e *InvalidCheckDateError) Code() string {
	return "INVALID_CHECK_DATE"
}

func (e *InvalidCheckDateError) IsUserError() bool {
	return true
}

type InvalidWeightError struct{}

func (e *InvalidWeightError) Error() string {
	return "weight must be greater than 0"
}

func (e *InvalidWeightError) Code() string {
	return "INVALID_WEIGHT"
}

func (e *InvalidWeightError) IsUserError() bool {
	return true
}

type InvalidPictureTypeError struct {
	PictureType string
}

func (e *InvalidPictureTypeError) Error() string {
	return fmt.Sprintf("invalid picture type: %s", e.PictureType)
}

func (e *InvalidPictureTypeError) Code() string {
	return "INVALID_PICTURE_TYPE"
}

func (e *InvalidPictureTypeError) IsUserError() bool {
	return true
}

type FileUploadError struct {
	Message string
}

func (e *FileUploadError) Error() string {
	return fmt.Sprintf("file upload error: %s", e.Message)
}

func (e *FileUploadError) Code() string {
	return "FILE_UPLOAD_ERROR"
}

func (e *FileUploadError) IsUserError() bool {
	return false
}

type WorkoutSessionNotFoundError struct {
	ID int
}

func (e *WorkoutSessionNotFoundError) Error() string {
	return fmt.Sprintf("workout session %d not found", e.ID)
}

func (e *WorkoutSessionNotFoundError) Code() string {
	return "WORKOUT_SESSION_NOT_FOUND"
}

func (e *WorkoutSessionNotFoundError) IsUserError() bool {
	return true
}

type LogExerciseSessionNotFoundError struct {
	ID int
}

func (e *LogExerciseSessionNotFoundError) Error() string {
	return fmt.Sprintf("exercise log %d not found", e.ID)
}

func (e *LogExerciseSessionNotFoundError) Code() string {
	return "EXERCISE_LOG_NOT_FOUND"
}

func (e *LogExerciseSessionNotFoundError) IsUserError() bool {
	return true
}

type InvalidSessionStatusTransitionError struct {
	From string
	To   string
}

func (e *InvalidSessionStatusTransitionError) Error() string {
	return fmt.Sprintf("cannot transition session from %s to %s", e.From, e.To)
}

func (e *InvalidSessionStatusTransitionError) Code() string {
	return "INVALID_SESSION_STATUS_TRANSITION"
}

func (e *InvalidSessionStatusTransitionError) IsUserError() bool {
	return true
}

type SessionAlreadyExistsError struct {
	RoutineID int
	Date      string
}

func (e *SessionAlreadyExistsError) Error() string {
	return fmt.Sprintf("session already exists for routine %d on %s", e.RoutineID, e.Date)
}

func (e *SessionAlreadyExistsError) Code() string {
	return "SESSION_ALREADY_EXISTS"
}

func (e *SessionAlreadyExistsError) IsUserError() bool {
	return true
}
