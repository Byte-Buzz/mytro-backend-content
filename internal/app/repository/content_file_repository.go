package repository

import (
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/database/dto"

	"github.com/google/uuid"
)

func filesToDomain(dtos []dto.ContentFileDTO) []models.ContentFile {
	if len(dtos) == 0 {
		return []models.ContentFile{}
	}

	files := make([]models.ContentFile, 0, len(dtos))
	for _, f := range dtos {
		contentID, err := uuid.Parse(f.ContentID)
		if err != nil {
			continue
		}

		fileID, err := uuid.Parse(f.FileID)
		if err != nil {
			continue
		}

		files = append(files, models.ContentFile{
			ContentID: contentID,
			FileID:    fileID,
			Role:      models.FileRole(f.Role),
			Width:     f.Width,
			Height:    f.Height,
		})
	}

	return files
}

func filesToDTO(files []models.ContentFile) []dto.ContentFileDTO {
	if len(files) == 0 {
		return []dto.ContentFileDTO{}
	}

	result := make([]dto.ContentFileDTO, 0, len(files))
	for _, f := range files {
		result = append(result, dto.ContentFileDTO{
			ContentID: f.ContentID.String(),
			FileID:    f.FileID.String(),
			Role:      string(f.Role),
			Width:     f.Width,
			Height:    f.Height,
		})
	}

	return result
}
