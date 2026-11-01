package models

import (
	"time"

	"github.com/google/uuid"
)

type Content struct {
	ID                 uuid.UUID
	OwnerID            uuid.UUID
	Type               ContentType
	Title              *string
	Description        *string
	Visibility         ContentVisibility
	Status             ContentStatus
	PublishRequestedAt *time.Time
	PublishedAt        *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Deleted            bool
	DeletedAt          *time.Time
	Metadata           *ContentMetadata
	Files              []ContentFile
}
