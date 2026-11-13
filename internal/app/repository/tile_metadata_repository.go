package repository

import (
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/infrastructure/database/dto"

	"github.com/google/uuid"
)

func metadataToDomain(dto *dto.TileMetadataDTO) *models.TileMetadata {
	if dto == nil {
		return nil
	}

	tileID, err := uuid.Parse(dto.TileID)
	if err != nil {
		return nil
	}

	return &models.TileMetadata{
		TileID:          tileID,
		DurationSeconds: dto.DurationSeconds,
		DominantColor:   dto.DominantColor,
	}
}

func metadataToDTO(m *models.TileMetadata) *dto.TileMetadataDTO {
	if m == nil {
		return nil
	}

	return &dto.TileMetadataDTO{
		TileID:          m.TileID.String(),
		DurationSeconds: m.DurationSeconds,
		DominantColor:   m.DominantColor,
	}
}
