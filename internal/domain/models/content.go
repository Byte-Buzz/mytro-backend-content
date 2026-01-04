package models

import "time"

type Content struct {
	ID          string
	OwnerID     string
	Type        ContentType
	Title       *string
	Description *string
	Visibility  ContentVisibility
	Status      ContentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
	DeletedAt   *time.Time
	Metadata    *ContentMetadata
	Files       []ContentFile
}
