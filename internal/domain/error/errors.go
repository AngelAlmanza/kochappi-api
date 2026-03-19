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
