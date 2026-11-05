package models

import "github.com/google/uuid"

type ContentMetadata struct {
	ContentID       uuid.UUID
	DurationSeconds *int
	DominantColor   *string
}
