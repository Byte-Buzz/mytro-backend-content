package service

import (
	"context"
	"errors"
	"mytro-backend-tile/internal/app/repository"
	"mytro-backend-tile/internal/domain/models"
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

type TileService struct {
	tileRepository repository.TileRepository
	fileRepository repository.FileRepository
	txManager      repository.TxManager
}

func NewTileService(tileRepository repository.TileRepository, fileRepository repository.FileRepository, txManager repository.TxManager) TileService {
	return TileService{
		tileRepository: tileRepository,
		fileRepository: fileRepository,
		txManager:      txManager,
	}
}

func (s TileService) CreateUpload(ctx context.Context, upload models.CreateUpload) (*models.CreatedUpload, error) {
	id := uuid.New()

	extension, exists := AvailableContentTypes[upload.FileType]
	if !exists {
		return nil, ErrUnsupportedContentType
	}

	filename := id.String() + extension
	isVideo := slices.Contains(VideoContentTypes, upload.FileType)

	err := s.tileRepository.Create(ctx, &models.Tile{
		ID:         id,
		OwnerID:    upload.OwnerID,
		Type:       getFileType(isVideo),
		Status:     models.TileStatusDraft,
		Visibility: models.TileVisibilityPrivate,
	})
	if err != nil {
		return nil, err
	}

	url, err := s.fileRepository.CreateUploadURL(ctx, models.UploadURLData{
		TileID:      id,
		Filename:    filename,
		MaxSize:     getMaxFile(isVideo),
		IsVideo:     isVideo,
		ContentType: upload.FileType,
		Visibility:  models.TileVisibilityPrivate,
	})
	if err != nil {
		return nil, err
	}

	createdURL := models.CreatedUpload{
		TileID:    id,
		FileID:    url.FileID,
		UploadURL: url.UploadURL,
		ExpiresAt: url.ExpiresAt,
	}

	return &createdURL, err
}

func (s TileService) CheckTileOwner(ctx context.Context, tileID uuid.UUID, userID uuid.UUID) (bool, error) {
	ownerID, err := s.tileRepository.GetOwnerByID(ctx, tileID)

	return ownerID == userID, err
}

func (s TileService) GetNewTiles(ctx context.Context, userID uuid.UUID) ([]models.Tile, error) {
	tileList, err := s.tileRepository.GetAllNewByOwner(ctx, userID)
	if err != nil {
		return nil, err
	}

	return tileList, nil
}

func getFileType(isVideo bool) models.TileType {
	if isVideo {
		return models.TileTypeVideo
	} else {
		return models.TileTypeImage
	}
}

func getMaxFile(isVideo bool) uint32 {
	if isVideo {
		return 200 * 1024 * 1024
	} else {
		return 20 * 1024 * 1024
	}
}
