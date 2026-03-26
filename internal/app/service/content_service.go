package service

import (
	"context"
	"mytro-backend-content/internal/app/repository"
	"mytro-backend-content/internal/domain/models"

	"github.com/google/uuid"
)

type ContentService struct {
	contentRepository repository.ContentRepository
	fileRepository    repository.FileRepository
}

func NewContentService(contentRepository repository.ContentRepository, fileRepository repository.FileRepository) ContentService {
	return ContentService{
		contentRepository: contentRepository,
		fileRepository:    fileRepository,
	}
}

func (s ContentService) CreateUpload(ctx context.Context) (*models.UploadURL, error) {
	content := models.Content{
		ID: uuid.New(),
	}

	return s.fileRepository.CreateUploadURL(ctx, models.UploadURLData{
		ContentID:   content.ID,
		Filename:    content.ID.String(),
		MaxSize:     20 * 1024 * 1024,
		IsVideo:     false,
		ContentType: "image/jpeg",
		Visibility:  models.ContentVisibilityPublic,
	})
}

func (s ContentService) CheckContentOwner(ctx context.Context, contentId uuid.UUID, ownerID uuid.UUID) (bool, error) {
	owner, err := s.contentRepository.GetOwnerById(ctx, contentId)

	return owner == ownerID, err
}

func (s ContentService) CompleteUpload(ctx context.Context, contentId uuid.UUID) error {
	// TODO: Check if the file is already uploaded
	return context.TODO().Err()
}

func (s ContentService) GetContentStatus(ctx context.Context, id uuid.UUID) (*models.Content, error) {
	return s.contentRepository.FindByID(ctx, id)
}
