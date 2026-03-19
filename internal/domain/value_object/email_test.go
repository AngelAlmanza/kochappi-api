package value_object

import "testing"

func TestNewEmail_ShouldAcceptValidEmails(t *testing.T) {
	validEmails := []string{
		"user@example.com",
		"USER@EXAMPLE.COM",
		"user.name@domain.co",
		"user+tag@example.com",
	}

	for _, email := range validEmails {
		e, err := NewEmail(email)
		if err != nil {
			t.Errorf("Expected email '%s' to be valid, got error: %v", email, err)
		}
		if e.String() == "" {
			t.Errorf("Expected non-empty email value for '%s'", email)
		}
	}
}

func TestNewEmail_ShouldNormalizeToLowercase(t *testing.T) {
	e, err := NewEmail("User@EXAMPLE.COM")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if e.String() != "user@example.com" {
		t.Errorf("Expected 'user@example.com', got '%s'", e.String())
	}
}

func TestNewEmail_ShouldTrimWhitespace(t *testing.T) {
	e, err := NewEmail("  user@example.com  ")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if e.String() != "user@example.com" {
		t.Errorf("Expected 'user@example.com', got '%s'", e.String())
	}
}

func TestNewEmail_ShouldRejectInvalidEmails(t *testing.T) {
	invalidEmails := []string{
		"",
		"not-an-email",
		"@example.com",
		"user@",
		"user@.com",
		"user@com",
	}

	for _, email := range invalidEmails {
		_, err := NewEmail(email)
		if err == nil {
			t.Errorf("Expected email '%s' to be invalid", email)
		}
	}
}
