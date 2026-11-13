package models

import (
	"time"

	"github.com/google/uuid"
)

type Tile struct {
	ID                 uuid.UUID
	OwnerID            uuid.UUID
	Type               TileType
	Title              *string
	Description        *string
	Visibility         TileVisibility
	Status             TileStatus
	PublishRequestedAt *time.Time
	PublishedAt        *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Deleted            bool
	DeletedAt          *time.Time
	Metadata           *TileMetadata
	Files              []TileFile
}
