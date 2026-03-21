package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LocalFileStorage struct {
	basePath string
	baseURL  string
}

func NewLocalFileStorage(basePath, baseURL string) *LocalFileStorage {
	return &LocalFileStorage{
		basePath: basePath,
		baseURL:  baseURL,
	}
}

func (s *LocalFileStorage) Upload(_ context.Context, filename string, file io.Reader) (string, error) {
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	filePath := filepath.Join(s.basePath, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	url := s.baseURL + "/" + filename
	return url, nil
}

func (s *LocalFileStorage) Delete(_ context.Context, url string) error {
	filename := strings.TrimPrefix(url, s.baseURL+"/")
	filePath := filepath.Join(s.basePath, filename)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
