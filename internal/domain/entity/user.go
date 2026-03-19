package entity

import "time"

type Role string

const (
	ROLE_TRAINER Role = "trainer"
	ROLE_CLIENT  Role = "client"
)

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Role         Role
	OTPCode      string
	OTPExpiresAt *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(id, name, email, passwordHash string, role Role) *User {
	now := time.Now()
	return &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (u *User) IsTrainer() bool {
	return u.Role == ROLE_TRAINER
}

func (u *User) IsClient() bool {
	return u.Role == ROLE_CLIENT
}

func (u *User) SetOTP(code string, expiresAt time.Time) {
	u.OTPCode = code
	u.OTPExpiresAt = &expiresAt
}

func (u *User) ClearOTP() {
	u.OTPCode = ""
	u.OTPExpiresAt = nil
}

func (u *User) IsOTPValid(code string) bool {
	if u.OTPCode == "" || u.OTPExpiresAt == nil {
		return false
	}
	if time.Now().After(*u.OTPExpiresAt) {
		return false
	}
	return u.OTPCode == code
}
