package repository

import (
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/database/dto"
)

func metadataToDomain(dto *dto.ContentMetadataDTO) *models.ContentMetadata {
	if dto == nil {
		return nil
	}

	return &models.ContentMetadata{
		ContentID:       dto.ContentID,
		DurationSeconds: dto.DurationSeconds,
		DominantColor:   dto.DominantColor,
	}
}

func metadataToDTO(m *models.ContentMetadata) *dto.ContentMetadataDTO {
	if m == nil {
		return nil
	}

	return &dto.ContentMetadataDTO{
		ContentID:       m.ContentID,
		DurationSeconds: m.DurationSeconds,
		DominantColor:   m.DominantColor,
	}
}
