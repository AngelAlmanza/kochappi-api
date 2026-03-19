package entity

import (
	"testing"
	"time"
)

func TestNewUser_ShouldCreateUserWithCorrectFields(t *testing.T) {
	user := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)

	if user.ID != "id-1" {
		t.Errorf("Expected ID 'id-1', got '%s'", user.ID)
	}
	if user.Name != "John" {
		t.Errorf("Expected Name 'John', got '%s'", user.Name)
	}
	if user.Email != "john@example.com" {
		t.Errorf("Expected Email 'john@example.com', got '%s'", user.Email)
	}
	if user.Role != ROLE_TRAINER {
		t.Errorf("Expected Role 'trainer', got '%s'", user.Role)
	}
	if user.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestUser_IsTrainer(t *testing.T) {
	trainer := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)
	client := NewUser("id-2", "Jane", "jane@example.com", "hashed", ROLE_CLIENT)

	if !trainer.IsTrainer() {
		t.Error("Expected trainer.IsTrainer() to be true")
	}
	if trainer.IsClient() {
		t.Error("Expected trainer.IsClient() to be false")
	}
	if !client.IsClient() {
		t.Error("Expected client.IsClient() to be true")
	}
	if client.IsTrainer() {
		t.Error("Expected client.IsTrainer() to be false")
	}
}

func TestUser_SetOTP_And_IsOTPValid(t *testing.T) {
	user := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)

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
	user := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)

	expiresAt := time.Now().Add(-1 * time.Minute)
	user.SetOTP("123456", expiresAt)

	if user.IsOTPValid("123456") {
		t.Error("Expected expired OTP to be invalid")
	}
}

func TestUser_IsOTPValid_ShouldReturnFalseWhenNoOTPSet(t *testing.T) {
	user := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)

	if user.IsOTPValid("123456") {
		t.Error("Expected OTP to be invalid when no OTP set")
	}
}

func TestUser_ClearOTP(t *testing.T) {
	user := NewUser("id-1", "John", "john@example.com", "hashed", ROLE_TRAINER)

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
