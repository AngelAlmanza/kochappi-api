package entity

import (
	"testing"
	"time"
)

func TestUser_SetOTP_And_IsOTPValid(t *testing.T) {
	user := NewUser("John", "john@example.com", "hashed", ROLE_TRAINER)

	expiresAt := time.Now().Add(10 * time.Minute)
	user.SetOTP("123456", expiresAt)

	if !user.IsOTPValid("123456") {
		t.Error("Expected OTP '123456' to be valid")
	}
	if user.IsOTPValid("654321") {
		t.Error("Expected OTP '654321' to be invalid")
	}
}

func TestUser_IsOTPValid_ShouldReturnFalseWhenExpired(t *testing.T) {
	user := NewUser("John", "john@example.com", "hashed", ROLE_TRAINER)

	expiresAt := time.Now().Add(-1 * time.Minute)
	user.SetOTP("123456", expiresAt)

	if user.IsOTPValid("123456") {
		t.Error("Expected expired OTP to be invalid")
	}
}

func TestUser_IsOTPValid_ShouldReturnFalseWhenNoOTPSet(t *testing.T) {
	user := NewUser("John", "john@example.com", "hashed", ROLE_TRAINER)

	if user.IsOTPValid("123456") {
		t.Error("Expected OTP to be invalid when no OTP set")
	}
}

func TestUser_ClearOTP(t *testing.T) {
	user := NewUser("John", "john@example.com", "hashed", ROLE_TRAINER)

	expiresAt := time.Now().Add(10 * time.Minute)
	user.SetOTP("123456", expiresAt)
	user.ClearOTP()

	if user.OTPCode != "" {
		t.Errorf("Expected OTPCode to be empty, got '%s'", user.OTPCode)
	}
	if user.OTPExpiresAt != nil {
		t.Error("Expected OTPExpiresAt to be nil")
	}
	if user.IsOTPValid("123456") {
		t.Error("Expected OTP to be invalid after clearing")
	}
}
