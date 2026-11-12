package repository

import (
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/infrastructure/database/dto"

	"github.com/google/uuid"
)

func metadataToDomain(dto *dto.TileMetadataDTO) *models.ContentMetadata {
	if dto == nil {
		return nil
	}

	contentID, err := uuid.Parse(dto.TileID)
	if err != nil {
		return nil
	}

	return &models.ContentMetadata{
		ContentID:       contentID,
		DurationSeconds: dto.DurationSeconds,
		DominantColor:   dto.DominantColor,
	}
}

func metadataToDTO(m *models.ContentMetadata) *dto.TileMetadataDTO {
	if m == nil {
		return nil
	}

	return &dto.TileMetadataDTO{
		TileID:          m.ContentID.String(),
		DurationSeconds: m.DurationSeconds,
		DominantColor:   m.DominantColor,
	}
}
