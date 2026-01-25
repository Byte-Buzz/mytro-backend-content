package repository

import (
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/database/dto"

	"github.com/google/uuid"
)

func metadataToDomain(dto *dto.ContentMetadataDTO) *models.ContentMetadata {
	if dto == nil {
		return nil
	}

	contentID, err := uuid.Parse(dto.ContentID)
	if err != nil {
		return nil
	}

	return &models.ContentMetadata{
		ContentID:       contentID,
		DurationSeconds: dto.DurationSeconds,
		DominantColor:   dto.DominantColor,
	}
}

func metadataToDTO(m *models.ContentMetadata) *dto.ContentMetadataDTO {
	if m == nil {
		return nil
	}

	return &dto.ContentMetadataDTO{
		ContentID:       m.ContentID.String(),
		DurationSeconds: m.DurationSeconds,
		DominantColor:   m.DominantColor,
	}
}
