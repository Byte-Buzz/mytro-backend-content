package models

import (
	"time"

	"github.com/google/uuid"
)

type UploadURLData struct {
	ContentID   uuid.UUID
	Filename    string
	MaxSize     uint32
	ContentType string
	Visibility  ContentVisibility
	IsVideo     bool
}

type UploadURL struct {
	ContentID uuid.UUID
	FileID    string
	UploadURL string
	ExpiresAt time.Time
}

type CreateUpload struct {
	FileType string
	OwnerID  uuid.UUID
}

type CreatedUpload struct {
	ContentID uuid.UUID
	FileID    string
	UploadURL string
	ExpiresAt time.Time
}
