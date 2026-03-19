package value_object

import "unicode/utf8"

const (
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 72
)

type Password struct {
	value string
}

func NewPassword(password string) (Password, error) {
	length := utf8.RuneCountInString(password)
	if length < MIN_PASSWORD_LENGTH {
		return Password{}, ErrPasswordTooShort
	}
	if length > MAX_PASSWORD_LENGTH {
		return Password{}, ErrPasswordTooLong
	}
	return Password{value: password}, nil
}

func (p Password) String() string {
	return p.value
}

var ErrPasswordTooShort = &ValidationError{
	Field:   "password",
	Message: "password must be at least 8 characters",
}

var ErrPasswordTooLong = &ValidationError{
	Field:   "password",
	Message: "password must be at most 72 characters",
}
