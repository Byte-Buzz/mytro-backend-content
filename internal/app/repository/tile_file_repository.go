package repository

import (
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/infrastructure/database/dto"

	"github.com/google/uuid"
)

func filesToDomain(dtos []dto.TileFileDTO) []models.TileFile {
	if len(dtos) == 0 {
		return []models.TileFile{}
	}

	files := make([]models.TileFile, 0, len(dtos))
	for _, f := range dtos {
		tileID, err := uuid.Parse(f.TileID)
		if err != nil {
			continue
		}

		fileID, err := uuid.Parse(f.FileID)
		if err != nil {
			continue
		}

		files = append(files, models.TileFile{
			TileID: tileID,
			FileID: fileID,
			Role:   models.FileRole(f.Role),
			Width:  f.Width,
			Height: f.Height,
		})
	}

	return files
}

func filesToDTO(files []models.TileFile) []dto.TileFileDTO {
	if len(files) == 0 {
		return []dto.TileFileDTO{}
	}

	result := make([]dto.TileFileDTO, 0, len(files))
	for _, f := range files {
		result = append(result, dto.TileFileDTO{
			TileID: f.TileID.String(),
			FileID: f.FileID.String(),
			Role:   string(f.Role),
			Width:  f.Width,
			Height: f.Height,
		})
	}

	return result
}
