package mock

import "context"

type MockRefreshTokenRepository struct {
	StoreFn             func(ctx context.Context, userID int, tokenID string, expiresAt int64) error
	ExistsFn            func(ctx context.Context, tokenID string) (bool, error)
	DeleteByIDFn        func(ctx context.Context, tokenID string) error
	DeleteAllByUserIDFn func(ctx context.Context, userID int) error
}

func (r *MockRefreshTokenRepository) Store(ctx context.Context, userID int, tokenID string, expiresAt int64) error {
	if r.StoreFn != nil {
		return r.StoreFn(ctx, userID, tokenID, expiresAt)
	}
	return nil
}

func (r *MockRefreshTokenRepository) Exists(ctx context.Context, tokenID string) (bool, error) {
	if r.ExistsFn != nil {
		return r.ExistsFn(ctx, tokenID)
	}
	return true, nil
}

func (r *MockRefreshTokenRepository) DeleteByID(ctx context.Context, tokenID string) error {
	if r.DeleteByIDFn != nil {
		return r.DeleteByIDFn(ctx, tokenID)
	}
	return nil
}

func (r *MockRefreshTokenRepository) DeleteAllByUserID(ctx context.Context, userID int) error {
	if r.DeleteAllByUserIDFn != nil {
		return r.DeleteAllByUserIDFn(ctx, userID)
	}
	return nil
}
