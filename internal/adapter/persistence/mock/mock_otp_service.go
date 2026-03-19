package mock

import "context"

type MockOTPService struct {
	GenerateCodeFn func() string
	SendFn         func(ctx context.Context, email string, code string) error
}

func (s *MockOTPService) GenerateCode() string {
	if s.GenerateCodeFn != nil {
		return s.GenerateCodeFn()
	}
	return "123456"
}

func (s *MockOTPService) Send(ctx context.Context, email string, code string) error {
	if s.SendFn != nil {
		return s.SendFn(ctx, email, code)
	}
	return nil
}
