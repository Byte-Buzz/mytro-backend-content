package models

import "github.com/google/uuid"

type UploadURLData struct {
	ContentID   uuid.UUID
	Filename    string
	MaxSize     int
	ContentType string
	Visibility  ContentVisibility
}
