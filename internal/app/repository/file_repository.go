package repository

import (
	"context"
	"mytro-backend-content/internal/domain/models"
)

type FileRepository interface {
	CreateUploadURL(ctx context.Context, data models.UploadURLData) (string, error)
}

type fileRepository struct {
}

// TODO: implement file repository
func NewFileRepository() FileRepository {
	return &fileRepository{}
}

func (r *fileRepository) CreateUploadURL(ctx context.Context, data models.UploadURLData) (string, error) {
	return "", nil
}
