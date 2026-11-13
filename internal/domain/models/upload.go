package models

import (
	"time"

	"github.com/google/uuid"
)

type UploadURLData struct {
	TileID      uuid.UUID
	Filename    string
	MaxSize     uint32
	ContentType string
	Visibility  TileVisibility
	IsVideo     bool
}

type UploadURL struct {
	FileID    string
	UploadURL string
	ExpiresAt time.Time
}

type CreateUpload struct {
	FileType string
	OwnerID  uuid.UUID
}

type CreatedUpload struct {
	TileID    uuid.UUID
	FileID    string
	UploadURL string
	ExpiresAt time.Time
}
