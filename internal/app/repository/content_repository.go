package repository

import (
	"context"
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/infrastructure/database/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContentRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.Content, error)
	GetOwnerByID(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	GetAllNewByOwner(ctx context.Context, id uuid.UUID) ([]models.Content, error)

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
	var content dto.TileDTO

	err := r.db.WithContext(ctx).
		Preload("Metadata").
		Preload("Files").
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

func (r *contentRepository) GetOwnerByID(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	var contentDTO dto.TileDTO

	result := r.db.WithContext(ctx).Select("owner_id").Where("id = ?", id.String()).Take(&contentDTO)

	if result.Error == gorm.ErrRecordNotFound {
		return uuid.Nil, nil
	}

	if result.Error != nil {
		return uuid.Nil, result.Error
	}

	return uuid.Parse(contentDTO.OwnerID)
}

func (r *contentRepository) GetAllNewByOwner(ctx context.Context, ownerID uuid.UUID) ([]models.Content, error) {
	contentDTOs := []dto.TileDTO{}

	result := r.db.WithContext(ctx).Where("owner_id = ? AND status <> 'draft' AND publish_requested_at IS NULL AND NOT deleted", ownerID.String()).Find(&contentDTOs)

	if result.Error == gorm.ErrRecordNotFound {
		return []models.Content{}, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	contents := make([]models.Content, 0, len(contentDTOs))

	for _, contentDTO := range contentDTOs {
		content, err := contentToDomain(&contentDTO)
		if err != nil {
			return []models.Content{}, err
		}

		contents = append(contents, *content)
	}

	return contents, nil
}

func (r *contentRepository) Create(ctx context.Context, content *models.Content) error {
	return r.db.WithContext(ctx).Create(contentToDTO(content)).Error
}

func contentToDomain(dto *dto.TileDTO) (*models.Content, error) {
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
		ID:                 id,
		OwnerID:            ownerID,
		Type:               models.ContentType(dto.Type),
		Title:              dto.Title,
		Description:        dto.Description,
		Visibility:         models.ContentVisibility(dto.Visibility),
		Status:             models.ContentStatus(dto.Status),
		PublishRequestedAt: dto.PublishRequestedAt,
		PublishedAt:        dto.PublishedAt,
		CreatedAt:          dto.CreatedAt,
		UpdatedAt:          dto.UpdatedAt,
		Deleted:            dto.Deleted,
		DeletedAt:          dto.DeletedAt,
		Metadata:           metadataToDomain(dto.Metadata),
		Files:              filesToDomain(dto.Files),
	}, nil
}

func contentToDTO(c *models.Content) *dto.TileDTO {
	if c == nil {
		return nil
	}

	return &dto.TileDTO{
		ID:                 c.ID.String(),
		OwnerID:            c.OwnerID.String(),
		Type:               string(c.Type),
		Title:              c.Title,
		Description:        c.Description,
		Visibility:         string(c.Visibility),
		Status:             string(c.Status),
		PublishRequestedAt: c.PublishRequestedAt,
		PublishedAt:        c.PublishedAt,
		CreatedAt:          c.CreatedAt,
		UpdatedAt:          c.UpdatedAt,
		Deleted:            c.Deleted,
		DeletedAt:          c.DeletedAt,
		Metadata:           metadataToDTO(c.Metadata),
		Files:              filesToDTO(c.Files),
	}
}
