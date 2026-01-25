package models

import "github.com/google/uuid"

type UploadURLData struct {
	ContentID   uuid.UUID
	Filename    string
	MaxSize     uint32
	ContentType string
	Visibility  ContentVisibility
	IsVideo     bool
}

type UploadURL struct {
	FileID    string
	UploadURL string
	ExpiresAt uint64
}
