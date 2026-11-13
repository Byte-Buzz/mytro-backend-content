package repository

import (
	"context"
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/infrastructure/database/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TileRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.Tile, error)
	GetOwnerByID(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	GetAllNewByOwner(ctx context.Context, id uuid.UUID) ([]models.Tile, error)

	Create(ctx context.Context, tile *models.Tile) error
}

type tileRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

// NewTileRepository creates a new TileRepository instance.
//
// The logger parameter is used to log messages from the repository.
// The db parameter is used to interact with the database.
//
// The returned TileRepository instance is safe for use by multiple
// goroutines.
func NewTileRepository(logger *zap.Logger, db *gorm.DB) TileRepository {
	return &tileRepository{
		logger: logger,
		db:     db,
	}
}

func (r *tileRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Tile, error) {
	var tile dto.TileDTO

	err := r.db.WithContext(ctx).
		Preload("Metadata").
		Preload("Files").
		Where("id = ?", id).
		First(&tile).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return tileToDomain(&tile)
}

func (r *tileRepository) GetOwnerByID(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	var tileDTO dto.TileDTO

	result := r.db.WithContext(ctx).Select("owner_id").Where("id = ?", id.String()).Take(&tileDTO)

	if result.Error == gorm.ErrRecordNotFound {
		return uuid.Nil, nil
	}

	if result.Error != nil {
		return uuid.Nil, result.Error
	}

	return uuid.Parse(tileDTO.OwnerID)
}

func (r *tileRepository) GetAllNewByOwner(ctx context.Context, ownerID uuid.UUID) ([]models.Tile, error) {
	tileDTOs := []dto.TileDTO{}

	result := r.db.WithContext(ctx).Where("owner_id = ? AND status <> 'draft' AND publish_requested_at IS NULL AND NOT deleted", ownerID.String()).Find(&tileDTOs)

	if result.Error == gorm.ErrRecordNotFound {
		return []models.Tile{}, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	tiles := make([]models.Tile, 0, len(tileDTOs))

	for _, tileDTO := range tileDTOs {
		tile, err := tileToDomain(&tileDTO)
		if err != nil {
			return []models.Tile{}, err
		}

		tiles = append(tiles, *tile)
	}

	return tiles, nil
}

func (r *tileRepository) Create(ctx context.Context, tile *models.Tile) error {
	return r.db.WithContext(ctx).Create(tileToDTO(tile)).Error
}

func tileToDomain(dto *dto.TileDTO) (*models.Tile, error) {
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

	return &models.Tile{
		ID:                 id,
		OwnerID:            ownerID,
		Type:               models.TileType(dto.Type),
		Title:              dto.Title,
		Description:        dto.Description,
		Visibility:         models.TileVisibility(dto.Visibility),
		Status:             models.TileStatus(dto.Status),
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

func tileToDTO(c *models.Tile) *dto.TileDTO {
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
