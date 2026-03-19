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
