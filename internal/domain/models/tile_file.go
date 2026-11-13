package models

import "github.com/google/uuid"

type TileFile struct {
	TileID uuid.UUID
	FileID uuid.UUID
	Role   FileRole
	Width  *int
	Height *int
}
