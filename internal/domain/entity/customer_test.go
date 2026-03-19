package entity

import (
	"testing"
	"time"
)

func TestCustomer_GetAge_ShouldReturnCorrectAge(t *testing.T) {
	birthdate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	customer := NewCustomer("John", birthdate, 1)

	expected := time.Now().Year() - 1990
	if customer.GetAge() != expected {
		t.Errorf("Expected age %d, got %d", expected, customer.GetAge())
	}
}

func TestCustomer_GetAge_ShouldReturn24BeforeBirthdayThisYear(t *testing.T) {
	now := time.Now()
	// Birthday set to December 31: before that date, the person is still 24.
	birthdate := time.Date(now.Year()-25, time.December, 31, 0, 0, 0, 0, time.UTC)
	customer := NewCustomer("Jane", birthdate, 2)

	expected := 25
	if now.Month() < time.December || (now.Month() == time.December && now.Day() < 31) {
		expected = 24
	}
	if customer.GetAge() != expected {
		t.Errorf("Expected age %d, got %d", expected, customer.GetAge())
	}
}
