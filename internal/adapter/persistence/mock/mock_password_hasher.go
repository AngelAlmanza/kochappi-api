package mock

import "errors"

type MockPasswordHasher struct {
	HashFn    func(password string) (string, error)
	CompareFn func(hashedPassword, password string) error
}

func (h *MockPasswordHasher) Hash(password string) (string, error) {
	if h.HashFn != nil {
		return h.HashFn(password)
	}
	return "hashed_" + password, nil
}

func (h *MockPasswordHasher) Compare(hashedPassword, password string) error {
	if h.CompareFn != nil {
		return h.CompareFn(hashedPassword, password)
	}
	if hashedPassword == "hashed_"+password {
		return nil
	}
	return errors.New("password mismatch")
}
