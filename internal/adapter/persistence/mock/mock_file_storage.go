package mock

import (
	"context"
	"io"
)

type MockFileStorage struct {
	UploadFn func(ctx context.Context, filename string, file io.Reader) (string, error)
	DeleteFn func(ctx context.Context, url string) error
}

func (s *MockFileStorage) Upload(ctx context.Context, filename string, file io.Reader) (string, error) {
	if s.UploadFn != nil {
		return s.UploadFn(ctx, filename, file)
	}
	return "", nil
}

func (s *MockFileStorage) Delete(ctx context.Context, url string) error {
	if s.DeleteFn != nil {
		return s.DeleteFn(ctx, url)
	}
	return nil
}
