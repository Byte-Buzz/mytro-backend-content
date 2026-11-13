package models

import "github.com/google/uuid"

type TileMetadata struct {
	TileID          uuid.UUID
	DurationSeconds *int
	DominantColor   *string
}
