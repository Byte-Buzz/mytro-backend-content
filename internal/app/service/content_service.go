package service

import (
	"context"
	"mytro-backend-content/internal/app/repository"
	"mytro-backend-content/internal/domain/models"

	"github.com/google/uuid"
)

type ContentService struct {
	ContentRepository repository.ContentRepository
}

func NewContentService(contentRepository repository.ContentRepository) ContentService {
	return ContentService{
		ContentRepository: contentRepository,
	}
}

func (s ContentService) FindByID(ctx context.Context, id uuid.UUID) (*models.Content, error) {
	return s.ContentRepository.FindByID(ctx, id)
}

func (s ContentService) Create(ctx context.Context, content *models.Content) error {
	return s.ContentRepository.Create(ctx, content)
}
