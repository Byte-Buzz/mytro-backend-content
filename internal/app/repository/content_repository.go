package repository

import (
	"context"
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/database/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContentRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.Content, error)

	Create(ctx context.Context, content *models.Content) error
}

type contentRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

// NewContentRepository creates a new ContentRepository instance.
//
// The logger parameter is used to log messages from the repository.
// The db parameter is used to interact with the database.
//
// The returned ContentRepository instance is safe for use by multiple
// goroutines.
func NewContentRepository(logger *zap.Logger, db *gorm.DB) ContentRepository {
	return &contentRepository{
		logger: logger,
		db:     db,
	}
}

func (r *contentRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Content, error) {
	var content dto.ContentDTO

	err := r.db.WithContext(ctx).
		Preload("ContentMetadata").
		Preload("ContentFiles").
		Where("id = ?", id).
		First(&content).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return contentToDomain(&content)
}

func (r *contentRepository) Create(ctx context.Context, content *models.Content) error {
	return r.db.WithContext(ctx).Create(contentToDTO(content)).Error
}

func contentToDomain(dto *dto.ContentDTO) (*models.Content, error) {
	if dto == nil {
		return nil, nil
	}

	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	ownerID, err := uuid.Parse(dto.OwnerID)
	if err != nil {
		return nil, err
	}

	return &models.Content{
		ID:          id,
		OwnerID:     ownerID,
		Type:        models.ContentType(dto.Type),
		Title:       dto.Title,
		Description: dto.Description,
		Visibility:  models.ContentVisibility(dto.Visibility),
		Status:      models.ContentStatus(dto.Status),
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		Deleted:     dto.Deleted,
		DeletedAt:   dto.DeletedAt,
		Metadata:    metadataToDomain(dto.Metadata),
		Files:       filesToDomain(dto.Files),
	}, nil
}

func contentToDTO(c *models.Content) *dto.ContentDTO {
	if c == nil {
		return nil
	}

	return &dto.ContentDTO{
		ID:          c.ID.String(),
		OwnerID:     c.OwnerID.String(),
		Type:        string(c.Type),
		Title:       c.Title,
		Description: c.Description,
		Visibility:  string(c.Visibility),
		Status:      string(c.Status),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		Deleted:     c.Deleted,
		DeletedAt:   c.DeletedAt,
		Metadata:    metadataToDTO(c.Metadata),
		Files:       filesToDTO(c.Files),
	}
}
