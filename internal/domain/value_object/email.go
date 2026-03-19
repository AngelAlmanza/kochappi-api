package value_object

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	normalized := strings.TrimSpace(strings.ToLower(email))
	if !emailRegex.MatchString(normalized) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: normalized}, nil
}

func (e Email) String() string {
	return e.value
}

var ErrInvalidEmail = &ValidationError{Field: "email", Message: "invalid email format"}
