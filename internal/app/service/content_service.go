package service

import (
	"context"
	"errors"
	"mytro-backend-content/internal/app/repository"
	"mytro-backend-content/internal/domain/models"
	"slices"

	"github.com/google/uuid"
)

var (
	ErrUnsupportedContentType = errors.New("unsupported content type")

	AvailableContentTypes = map[string]string{
		"image/bmp":       ".bmp",
		"image/gif":       ".gif",
		"image/png":       ".png",
		"image/jpeg":      ".jpg",
		"image/tiff":      ".tif",
		"image/webp":      ".webp",
		"video/mp4":       ".mp4",
		"video/x-m4v":     ".m4v",
		"video/quicktime": ".mov",
	}
	VideoContentTypes = []string{"video/mp4", "video/x-m4v", "video/quicktime"}
)

type ContentService struct {
	contentRepository repository.ContentRepository
	fileRepository    repository.FileRepository
	txManager         repository.TxManager
}

func NewContentService(contentRepository repository.ContentRepository, fileRepository repository.FileRepository, txManager repository.TxManager) ContentService {
	return ContentService{
		contentRepository: contentRepository,
		fileRepository:    fileRepository,
		txManager:         txManager,
	}
}

func (s ContentService) CreateUpload(ctx context.Context, content models.CreateUpload) (*models.UploadURL, error) {
	id := uuid.New()

	extension, exists := AvailableContentTypes[content.FileType]
	if !exists {
		return nil, ErrUnsupportedContentType
	}

	filename := id.String() + extension
	isVideo := slices.Contains(VideoContentTypes, extension)

	s.contentRepository.Create(ctx, &models.Content{
		ID:         id,
		OwnerID:    content.OwnerID,
		Type:       getFileType(isVideo),
		Status:     models.ContentStatusDraft,
		Visibility: models.ContentVisibilityPrivate,
	})

	return s.fileRepository.CreateUploadURL(ctx, models.UploadURLData{
		ContentID:   id,
		Filename:    filename,
		MaxSize:     getMaxFile(isVideo),
		IsVideo:     isVideo,
		ContentType: content.FileType,
		Visibility:  models.ContentVisibilityPrivate,
	})
}

func (s ContentService) CheckContentOwner(ctx context.Context, contentId uuid.UUID, userID uuid.UUID) (bool, error) {
	ownerID, err := s.contentRepository.GetOwnerById(ctx, contentId)

	return ownerID == userID, err
}

func (s ContentService) GetNewContentList(ctx context.Context, userId uuid.UUID) ([]models.Content, error) {
	contentList, err := s.contentRepository.GetAllNewByOwner(ctx, userId)
	if err != nil {
		return nil, err
	}

	return contentList, nil
}

func getFileType(isVideo bool) models.ContentType {
	if isVideo {
		return models.ContentTypeVideo
	} else {
		return models.ContentTypeImage
	}
}

func getMaxFile(isVideo bool) uint32 {
	if isVideo {
		return 200 * 1024 * 1024
	} else {
		return 20 * 1024 * 1024
	}
}
