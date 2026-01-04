package repository

import (
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/database/dto"
)

func filesToDomain(dtos []dto.ContentFileDTO) []models.ContentFile {
	if len(dtos) == 0 {
		return []models.ContentFile{}
	}

	files := make([]models.ContentFile, 0, len(dtos))
	for _, f := range dtos {
		files = append(files, models.ContentFile{
			ContentID: f.ContentID,
			FileID:    f.FileID,
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
			ContentID: f.ContentID,
			FileID:    f.FileID,
			Role:      string(f.Role),
			Width:     f.Width,
			Height:    f.Height,
		})
	}

	return result
}
