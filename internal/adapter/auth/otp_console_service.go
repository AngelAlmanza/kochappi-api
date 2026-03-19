package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

type ConsoleOTPService struct{}

func NewConsoleOTPService() *ConsoleOTPService {
	return &ConsoleOTPService{}
}

func (s *ConsoleOTPService) GenerateCode() string {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "000000"
	}
	return fmt.Sprintf("%06d", n.Int64())
}

func (s *ConsoleOTPService) Send(ctx context.Context, email string, code string) error {
	fmt.Println("══════════════════════════════════════════")
	fmt.Printf("  OTP Code for %s\n", email)
	fmt.Printf("  Code: %s\n", code)
	fmt.Println("══════════════════════════════════════════")
	return nil
}
