package value_object

import (
	"strings"
	"testing"
)

func TestNewPassword_ShouldAcceptValidPasswords(t *testing.T) {
	validPasswords := []string{
		"12345678",
		"a_valid_password",
		"MyP@ssw0rd!",
	}

	for _, password := range validPasswords {
		p, err := NewPassword(password)
		if err != nil {
			t.Errorf("Expected password '%s' to be valid, got error: %v", password, err)
		}
		if p.String() != password {
			t.Errorf("Expected '%s', got '%s'", password, p.String())
		}
	}
}

func TestNewPassword_ShouldRejectTooShort(t *testing.T) {
	_, err := NewPassword("1234567")
	if err == nil {
		t.Error("Expected error for password shorter than 8 characters")
	}
	if err != ErrPasswordTooShort {
		t.Errorf("Expected ErrPasswordTooShort, got %v", err)
	}
}

func TestNewPassword_ShouldRejectTooLong(t *testing.T) {
	longPassword := strings.Repeat("a", 73)
	_, err := NewPassword(longPassword)
	if err == nil {
		t.Error("Expected error for password longer than 72 characters")
	}
	if err != ErrPasswordTooLong {
		t.Errorf("Expected ErrPasswordTooLong, got %v", err)
	}
}

func TestNewPassword_ShouldRejectEmptyString(t *testing.T) {
	_, err := NewPassword("")
	if err == nil {
		t.Error("Expected error for empty password")
	}
}

func TestNewPassword_ShouldAcceptExactlyMinLength(t *testing.T) {
	_, err := NewPassword("12345678")
	if err != nil {
		t.Errorf("Expected no error for 8 character password, got %v", err)
	}
}

func TestNewPassword_ShouldAcceptExactlyMaxLength(t *testing.T) {
	maxPassword := strings.Repeat("a", 72)
	_, err := NewPassword(maxPassword)
	if err != nil {
		t.Errorf("Expected no error for 72 character password, got %v", err)
	}
}
